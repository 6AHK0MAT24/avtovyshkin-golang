-- Migration: 004_insert_random_drivers.sql
-- Description: Insert 5 random drivers with all fields filled

INSERT INTO drivers (
    fldfirstname, fldlastname, fldmiddlename, fldphone, fldemail, fldbirthdate,
    flddriverlicensenumber, flddriverlicenseissuedate, flddriverlicenseexpirydate,
    flddriverlicensephoto, flddriverlicensescan, fldpassportseries, fldpassportnumber,
    fldpassportissuedate, fldpassportphoto, fldpassportscan, fldaddress,
    fldexperienceyears, fldstatus, fldphoto
) VALUES
(
    'Иван', 'Иванов', 'Иванович', '+79001234567', 'ivanov.ivan@example.com', '1985-03-15',
    '1234567890', '2015-06-20', '2025-06-20',
    '/uploads/drivers/license/ivanov_license.jpg', '/uploads/drivers/license/ivanov_license_scan.jpg',
    '4501', '123456', '2015-05-10',
    '/uploads/drivers/passport/ivanov_passport.jpg', '/uploads/drivers/passport/ivanov_passport_scan.jpg',
    'г. Москва, ул. Ленина, д. 10, кв. 25', 8, 'active', '/uploads/drivers/photo/ivanov.jpg'
),
(
    'Петр', 'Петров', 'Петрович', '+79002345678', 'petrov.petr@example.com', '1990-07-22',
    '2345678901', '2018-08-15', '2028-08-15',
    '/uploads/drivers/license/petrov_license.jpg', '/uploads/drivers/license/petrov_license_scan.jpg',
    '4502', '234567', '2018-07-05',
    '/uploads/drivers/passport/petrov_passport.jpg', '/uploads/drivers/passport/petrov_passport_scan.jpg',
    'г. Санкт-Петербург, пр. Невский, д. 15, кв. 30', 5, 'active', '/uploads/drivers/photo/petrov.jpg'
),
(
    'Сергей', 'Сидоров', 'Сергеевич', '+79003456789', 'sidorov.sergey@example.com', '1988-11-30',
    '3456789012', '2016-04-10', '2026-04-10',
    '/uploads/drivers/license/sidorov_license.jpg', '/uploads/drivers/license/sidorov_license_scan.jpg',
    '4503', '345678', '2016-03-20',
    '/uploads/drivers/passport/sidorov_passport.jpg', '/uploads/drivers/passport/sidorov_passport_scan.jpg',
    'г. Новосибирск, ул. Красный проспект, д. 20, кв. 15', 7, 'active', '/uploads/drivers/photo/sidorov.jpg'
),
(
    'Александр', 'Козлов', 'Александрович', '+79004567890', 'kozlov.alexander@example.com', '1992-02-14',
    '4567890123', '2019-09-25', '2029-09-25',
    '/uploads/drivers/license/kozlov_license.jpg', '/uploads/drivers/license/kozlov_license_scan.jpg',
    '4504', '456789', '2019-08-15',
    '/uploads/drivers/passport/kozlov_passport.jpg', '/uploads/drivers/passport/kozlov_passport_scan.jpg',
    'г. Екатеринбург, ул. Вайнера, д. 5, кв. 40', 4, 'active', '/uploads/drivers/photo/kozlov.jpg'
),
(
    'Дмитрий', 'Морозов', 'Дмитриевич', '+79005678901', 'morozov.dmitry@example.com', '1987-09-08',
    '5678901234', '2017-12-01', '2027-12-01',
    '/uploads/drivers/license/morozov_license.jpg', '/uploads/drivers/license/morozov_license_scan.jpg',
    '4505', '567890', '2017-11-10',
    '/uploads/drivers/passport/morozov_passport.jpg', '/uploads/drivers/passport/morozov_passport_scan.jpg',
    'г. Казань, ул. Баумана, д. 25, кв. 10', 6, 'active', '/uploads/drivers/photo/morozov.jpg'
);
