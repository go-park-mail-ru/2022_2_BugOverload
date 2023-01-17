-- Microservice user (while in the api microservice) - all user activity in app
CREATE USER user_app WITH LOGIN PASSWORD 'pass_microservice_2022_user';

-- Microservice auth
CREATE USER auth_app WITH LOGIN PASSWORD 'pass_microservice_2022_auth';

-- Microservice warehouse - all public content in app
CREATE USER warehouse_app WITH LOGIN PASSWORD 'pass_microservice_2022_warehouse';

-- Microservice image - for control image workflow
CREATE USER image_app WITH LOGIN PASSWORD 'pass_microservice_2022_image';
