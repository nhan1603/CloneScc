INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('IamID',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1000,'IamID','PremisesID','2023-07','2023-05','error','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1000,'test_account','test_account@gmail.com','test',
  '','','','','','',0,0,'',0,0,0,0,'',0,'0', 0,0, 'not eligible', '');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1000, '', false, 'Energy consumption not available');

-------
  
INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('auth0|60e5c63537b26948177cf800',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1001,'auth0|60e5c63537b26948177cf800','160312345125','2023-07','2023-05','ready','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1001,'test_account','test_account@gmail.com','test',
  'GOOD','GOOD','GOOD','15%','15%','20%',10.5,0.45,'GOOD',3.5,1.7,7.2,1.5,'AVERAGE',77,'VERY GOOD', 15.5,0.35, 'eligible', 'kCo2');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1001, 'https://google.com', true, '');

-------

INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('auth0|6008d0459db3352217337a0c',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1002,'auth0|6008d0459db3352217337a0c','497245123','2023-07','2023-05','ready','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1002,'test_account','test_account@gmail.com','test',
  'GOOD','GOOD','GOOD','15%','15%','20%',10.5,0.45,'GOOD',3.5,1.7,7.2,1.5,'AVERAGE',77,'VERY GOOD', 15.5,0.35, 'eligible', 'kCo2');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1002, 'https://google.com', true, '');

-------

INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('auth0|6267663421fcc63fbb39f222',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1003,'auth0|6267663421fcc63fbb39f222','553167489','2023-07','2023-05','ready','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1003,'test_account','test_account@gmail.com','test',
  'GOOD','GOOD','GOOD','15%','15%','20%',10.5,0.45,'GOOD',3.5,1.7,7.2,1.5,'AVERAGE',77,'VERY GOOD', 15.5,0.35, 'eligible', 'kCo2');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1003, 'https://google.com', true, '');
  
-------

INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('auth0|60e5c43d37b26948177cf7c4',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1004,'auth0|60e5c43d37b26948177cf7c4','4050573424','2023-07','2023-05','error','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1004,'test_account','test_account@gmail.com','test',
  '','','','','','',0,0,'',0,0,0,0,'',0,'0', 0,0, 'not eligible', '');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1004, '', false, 'Energy consumption not available');
  
-------

INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('auth0|60b735bb0af2a7586d5d272b',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1005,'auth0|60b735bb0af2a7586d5d272b','091342443241','2023-07','2023-05','error','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1005,'test_account','test_account@gmail.com','test',
  '','','','','','',0,0,'',0,0,0,0,'',0,'0', 0,0, 'not eligible', '');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1005, '', false, 'Energy consumption not available');
  
-------

INSERT INTO accounts(id,total_points,created_at,created_by,updated_at,updated_by) 
VALUES 
('auth0|62c2a18b6a4ae170a3fcb228',0,E'2021-04-22 00:00:00','test_account',E'2021-04-22 00:00:00','admin');

INSERT INTO green_up_report_request(id,account_id,premises_id,request_month_year,report_month_year,client_status,processing_status)
VALUES
(1006,'auth0|62c2a18b6a4ae170a3fcb228','945321123','2023-07','2023-05','error','initial');

INSERT INTO green_up_report_data
(
  report_id,user_name, user_email, premise_address,
  elec_consumption, water_consumption, gas_consumption, elec_reduction, water_reduction, gas_reduction, total_co2, co2_reduction, co2_survey_banding, food_co2, commute_co2, spending_co2, travel_co2, greenup_banding, green_score, green_banding ,utility_carbon_emission, utility_carbon_emission_pct_change, eligibility, unit)
VALUES(1006,'test_account','test_account@gmail.com','test',
  '','','','','','',0,0,'',0,0,0,0,'',0,'0', 0,0, 'not eligible', '');

INSERT INTO public.green_up_report_result(
	report_id, url, eligible, error)
	VALUES (1006, '', false, 'Energy consumption not available');
  
-------