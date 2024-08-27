INSERT INTO users(id, display_name, email, password, user_role)
VALUES (100, 'John', 'john@scc.com', '$2a$14$HGdzOYMZu81gzCxAE8iSz.BIjNMivFznMff0cbBw7Au2UJby13YCK', 'OPERATION_USER'),
       (101, 'Thomas', 'thomas@scc.com', '$2a$14$KZBv50VSGBpIfmtZ8a4G4euHuql5F2JShgCIXlxlrFUfkEotOU8b2', 'SECURITY_GUARD'),
       (102, 'William', 'william@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD');

-- Dummy data for Premises
INSERT INTO premises (id, name, location, premises_code, description, cctv_count) VALUES
(100, 'Sunrise Tower', '307/12 Nguyen Van Troi St, W1, Tan Binh', 'P100', 'Sunrise Tower', 4),
(101, 'Bitexco Financial Tower', '2 Hai Ba Trung St, Ben Nghe, District 1', 'P101', 'Bitexco Financial Tower', 4);

-- Dummy data for CCTV Devices
INSERT INTO cctv_devices (id, premise_id, device_name, device_code, is_active, floor_number) VALUES
(200, 100, 'CCTV 1', 'cctv_1', true, 1),
(201, 100, 'CCTV 2', 'cctv_2', true, 1),
(202, 101, 'CCTV 3', 'cctv_3', true, 2);


-- Dummy data for alerts
INSERT INTO alerts (id, cctv_device_id, type, description, media_data, is_acknowledged, incident_at)
VALUES
    (300, 200, 'Suspicious Activities', 'TEST', '[]', TRUE, '2023-07-25 00:00:00.000000 +00:00'),
    (301, 200, 'Unauthorized Access', 'TEST', '[]', TRUE, '2023-07-26 00:00:00.000000 +00:00'),
    (302, 200, 'Suspicious Activities', 'TEST', '[]', TRUE, '2023-07-27 00:00:00.000000 +00:00'),
    (303, 201, 'Unauthorized Access', 'TEST', '[]', TRUE, '2023-07-29 00:00:00.000000 +00:00'),
    (304, 201, 'Suspicious Activities', 'TEST', '[]', TRUE, '2023-07-29 00:00:00.000000 +00:00'),
    (305, 202, 'Property Damage', 'TEST', '[]', TRUE, '2023-07-29 00:00:00.000000 +00:00'),
    (306, 202, 'Suspicious Activities', 'TEST', '[]', TRUE, '2023-07-29 00:00:00.000000 +00:00');

-- Dummy data for request
INSERT INTO verification_requests (id, alert_id, request_by, assigned_user_id, status, message, start_time, end_time)
VALUES
    (400, 300, 100, 101, 'NEW', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-25 00:00:00.000000 +00:00', null),
    (401, 300, 100, 101, 'NEW', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-26 00:00:00.000000 +00:00', null),
    (402, 300, 100, 102, 'RESOLVED', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-27 00:00:00.000000 +00:00', '2023-07-30 00:00:00.000000 +00:00'),
    (403, 301, 100, 102, 'NEW', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-27 00:00:00.000000 +00:00', null),
    (404, 301, 100, 101, 'NEW', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-27 00:00:00.000000 +00:00', null),
    (405, 302, 100, 101, 'RESOLVED', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-27 00:00:00.000000 +00:00', '2023-07-29 00:00:00.000000 +00:00'),
    (406, 303, 100, 102, 'NEW', 'Property Damage Detected! Please check the monitored area and take screenshots for verification. Report any damage or suspicious activities to the authorities.', '2023-07-29 00:00:00.000000 +00:00', null);

-- Dummy data for request responses
INSERT INTO verification_request_responses (id, verification_request_id, message, media_data, verified_at)
VALUES
    (500, 400, 'Ive checked the monitored area, and it''s a false alarm. No property damage or suspicious activities were found.', '[]', '2023-07-25 00:00:00.000000 +00:00'),
    (501, 400, 'Ive checked the monitored area, and it''s a false alarm. No property damage or suspicious activities were found.', '[]', '2023-07-26 00:00:00.000000 +00:00'),
    (502, 401, 'Ive checked the monitored area, and it''s a false alarm. No property damage or suspicious activities were found.', '[]', '2023-07-27 00:00:00.000000 +00:00'),
    (503, 402, 'Ive checked the monitored area, and it''s a false alarm. No property damage or suspicious activities were found.', '[]', '2023-07-30 00:00:00.000000 +00:00'),
    (504, 403, 'Ive checked the monitored area, and it''s a false alarm. No property damage or suspicious activities were found.', '[]', '2023-07-30 00:00:00.000000 +00:00');

