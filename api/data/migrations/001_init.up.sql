CREATE TABLE IF NOT EXISTS public.users
(
    id              BIGINT PRIMARY KEY,
    display_name    TEXT                     NOT NULL CONSTRAINT users_display_name_check CHECK (display_name <> ''::TEXT),
    email           TEXT                     NOT NULL CONSTRAINT users_email_check CHECK (email <> ''::TEXT),
    password        TEXT                     NOT NULL CONSTRAINT users_password_check CHECK (password <> ''::TEXT),
    user_role       TEXT                     NOT NULL CONSTRAINT users_user_role_check CHECK (user_role <> ''::TEXT),
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    UNIQUE (display_name),
    UNIQUE (email)
);
CREATE INDEX IF NOT EXISTS users_email_index ON users(email);

CREATE TABLE IF NOT EXISTS public.premises
(
    id              BIGINT PRIMARY KEY,
    name            TEXT                     NOT NULL CONSTRAINT premises_premises_name_check CHECK (name <> ''::TEXT),
    location        TEXT                     NOT NULL CONSTRAINT premises_premises_location_check CHECK (location <> ''::TEXT),
    premises_code   TEXT                     NOT NULL CONSTRAINT premises_premises_code_check CHECK (premises_code <> ''::TEXT),
    description     TEXT                     NOT NULL CONSTRAINT premises_description_check CHECK (description <> ''::TEXT),
    cctv_count      INT                      NOT NULL CONSTRAINT premises_cctv_count_check CHECK (cctv_count > 0::INT),
    created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    UNIQUE (premises_code)
);
CREATE INDEX IF NOT EXISTS premises_premises_code_index ON premises(premises_code);

CREATE TABLE IF NOT EXISTS public.cctv_devices
(
	id              BIGINT PRIMARY KEY,
    premise_id      BIGINT                   NOT NULL CONSTRAINT cctv_devices_premise_id_fkey REFERENCES premises (id),
	device_name     TEXT                     NOT NULL CONSTRAINT cctv_devices_device_name_check CHECK (device_name <> ''::TEXT),
	device_code     TEXT                     NOT NULL CONSTRAINT cctv_devices_device_code_check CHECK (device_code <> ''::TEXT),
    is_active       BOOLEAN                  NOT NULL DEFAULT FALSE,
    floor_number    INT,
	created_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
	updated_at      TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    UNIQUE (device_code)
);
CREATE INDEX IF NOT EXISTS cctv_devices_device_code_index ON cctv_devices(device_code);

CREATE TABLE IF NOT EXISTS public.alerts
(
    id               BIGINT PRIMARY KEY,
    cctv_device_id   BIGINT                   NOT NULL CONSTRAINT alerts_cctv_device_id_fkey REFERENCES cctv_devices (id),
    type             TEXT                     NOT NULL CONSTRAINT alerts_type_check CHECK (type <> ''::TEXT),
    description      TEXT                     NOT NULL CONSTRAINT alerts_description_check CHECK (description <> ''::TEXT),
    media_data       JSONB                    NOT NULL CONSTRAINT alerts_media_data_check CHECK (media_data <> '{}'::JSONB),
    is_acknowledged  BOOLEAN                  NOT NULL DEFAULT FALSE,
    incident_at      TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at       TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);
CREATE INDEX IF NOT EXISTS alerts_type_index ON alerts(type);

CREATE TABLE IF NOT EXISTS public.verification_requests(
    id                  BIGINT PRIMARY KEY,
    alert_id            BIGINT                   NOT NULL CONSTRAINT verification_requests_alert_id_fkey REFERENCES alerts (id),
    request_by          BIGINT                   NOT NULL CONSTRAINT verification_requests_request_by_fkey REFERENCES users (id),
    assigned_user_id    BIGINT                   NOT NULL CONSTRAINT verification_requests_sg_groups_assigned_user_id_fkey REFERENCES users (id),
    status              TEXT                     NOT NULL CONSTRAINT verification_requests_status_check CHECK (status <> ''::TEXT),
    message             TEXT,
    start_time          TIMESTAMP WITH TIME ZONE NOT NULL,
    end_time            TIMESTAMP WITH TIME ZONE,
    created_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.verification_request_responses(
    id                                  BIGINT PRIMARY KEY,
    verification_request_id             BIGINT                   NOT NULL CONSTRAINT verification_request_responses_verification_request_id_fkey REFERENCES verification_requests (id),
    message                             TEXT                     NOT NULL CONSTRAINT verification_request_responses_messgae_check CHECK (message <> ''::TEXT),
    media_data                          JSONB,
    verified_at                         TIMESTAMP WITH TIME ZONE NOT NULL,
    created_at                          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at                          TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS public.device_tokens
(
    id           BIGINT PRIMARY KEY,
    user_id      BIGINT                   NOT NULL CONSTRAINT device_user_id_fkey REFERENCES users (id),
    device_token TEXT                     NOT NULL CONSTRAINT device_token_check CHECK (device_token <> ''::TEXT),
    platform     TEXT                     NOT NULL CONSTRAINT device_platform_check CHECK (platform <> ''::TEXT),
    created_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    UNIQUE (user_id)
);
