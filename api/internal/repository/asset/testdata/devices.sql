INSERT INTO public.premises (id, name, location, premises_code, description, cctv_count) VALUES
(50, 'ABC Tower', 'HCMC', 'P00501', 'ABC Tower', 4);

INSERT INTO public.cctv_devices (id, premise_id, device_name, device_code, is_active, floor_number) VALUES
(100, 50, 'Camera 50', 'cctv_cam50', true, 1),
(101, 50, 'Camera 51', 'cctv_cam51', true, 2);
