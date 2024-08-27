INSERT INTO users(id, display_name, email, password, user_role) VALUES
(1, 'Operation', 'operation@scc.com', '$2a$14$h5xx6YnrxlmewN8MkrXpjeN34TVgxJXcx8k70goLZFZ/RcKmiRXgO', 'OPERATION_USER'),
(2, 'Security', 'security@scc.com', '$2a$14$yxTbtzimrMVmZPxxptK8H.2h5Q9q2agJ72Y7JvMH8xbjCx/BNe2xi', 'SECURITY_GUARD'),

-- OPERATION_USER 
(20, 'Admin', 'admin@scc.com', '$2a$14$HGdzOYMZu81gzCxAE8iSz.BIjNMivFznMff0cbBw7Au2UJby13YCK', 'OPERATION_USER'),
(21, 'Ky', 'ky@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),
(22, 'Hue', 'hue@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),
(23, 'Mai', 'mai@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),
(24, 'Vinh', 'vinh@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),
(25, 'Hai', 'hai@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),
(26, 'Tung', 'tung@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),
(27, 'Tin', 'tin@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'OPERATION_USER'),

--  SECURITY_GUARD
(52, 'Phuong SG', 'phuongsg@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD'),
(53, 'Don SG', 'donsg@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD'),
(54, 'My SG', 'mysg@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD'),
(55, 'Mai SG', 'maisg@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD'),
(56, 'Ky SG', 'kysg@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD'),
(57, 'Tin SG', 'tinsg@scc.com', '$2a$14$WUPB.LddBw6CyJ3bNIBfTO6OdTqWQ14lkL3xlXGs8si8YHY6ozhXC', 'SECURITY_GUARD');

-- Dummy data for Premises
INSERT INTO premises (id, name, location, premises_code, description, cctv_count)
VALUES
  (5, 'Sunrise Tower', '307/12 Nguyen Van Troi St, W1, Tan Binh', 'P001', 'Sunrise Tower', 4),
  (6, 'Bitexco Financial Tower', '2 Hai Ba Trung St, Ben Nghe, District 1', 'P002', 'Bitexco Financial Tower', 2),
  (7, 'Lua Office', '40 Nguyen Van Troi st., W1, Tan Binh, HCMC', 'P003', 'Lua Office', 7),
  (8, 'Head Office', '10 Truong Dinh st., W8, District 3, HCMC', 'P004', 'Head Office', 10);

-- Dummy data for CCTV Devices
INSERT INTO cctv_devices (id, premise_id, device_name, device_code, is_active, floor_number)
VALUES
    -- Sunrise Tower
    (20, 5, 'Camera 1', 'cctv_cam1', true, 1),
    (21, 5, 'Camera 2', 'cctv_cam2', true, 1),
    (22, 5, 'Camera 3', 'cctv_cam3', true, 2),
    (23, 5, 'Camera 4', 'cctv_cam4', true, 3),
    -- Bitexco Financial Tower
    (24, 6, 'Camera 5', 'cctv_cam5', true, 1),
    (25, 6, 'Camera 6', 'cctv_cam6', true, 2),
    -- Lua Office
    (30, 7, 'Camera 1', 'cctv_cam30', true, 0),
    (31, 7, 'Camera 2', 'cctv_cam31', true, 1),
    (32, 7, 'Camera 3', 'cctv_cam32', true, 1),
    (33, 7, 'Camera 4', 'cctv_cam33', true, 2),
    (34, 7, 'Camera 5', 'cctv_cam34', true, 2),
    (35, 7, 'Camera 6', 'cctv_cam35', true, 3),
    (36, 7, 'Camera 7', 'cctv_cam36', true, 4),
    -- Head Office
    (40, 8, 'CCTV 1', 'cctv_cam40', true, 0),
    (41, 8, 'CCTV 2', 'cctv_cam41', true, 1),
    (42, 8, 'CCTV 3', 'cctv_cam42', true, 1),
    (43, 8, 'CCTV 4', 'cctv_cam43', true, 2),
    (44, 8, 'CCTV 5', 'cctv_cam44', true, 2),
    (45, 8, 'CCTV 6', 'cctv_cam45', true, 3),
    (46, 8, 'CCTV 7', 'cctv_cam46', true, 4),
    (47, 8, 'CCTV 8', 'cctv_cam47', true, 5),
    (48, 8, 'CCTV 9', 'cctv_cam48', true, 6),
    (49, 8, 'CCTV 10', 'cctv_cam49', true, 7);
    