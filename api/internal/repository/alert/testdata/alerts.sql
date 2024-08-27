-- Dummy data for Premises
INSERT INTO premises (id, name, location, premises_code, description, cctv_count) VALUES
      (100, 'Sunrise Tower', '307/12 Nguyen Van Troi St, W1, Tan Binh', 'P100', 'Sunrise Tower', 4);

-- Dummy data for CCTV Devices
INSERT INTO cctv_devices (id, premise_id, device_name, device_code, is_active, floor_number) VALUES
     (200, 100, 'CCTV 1', 'cctv_1', true, 1);


-- Dummy data for alerts
INSERT INTO alerts (id, cctv_device_id, type, description, media_data, is_acknowledged, incident_at)
VALUES
    (300, 200, 'Suspicious Activities', 'TEST', '[]', TRUE, '2023-07-25 00:00:00.000000 +00:00'),
    (301, 200, 'Unauthorized Access', 'TEST', '[]', TRUE, '2023-07-26 00:00:00.000000 +00:00'),
    (302, 200, 'Suspicious Activities', 'TEST', '[]', TRUE, '2023-07-27 00:00:00.000000 +00:00');
