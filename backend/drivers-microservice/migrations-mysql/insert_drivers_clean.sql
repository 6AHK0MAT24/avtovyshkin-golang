-- Очистка таблицы перед вставкой данных
TRUNCATE TABLE drivers;

-- Вставка начальных данных водителей (MySQL версия)
INSERT INTO drivers (
    fldid, fldfirstname, fldlastname, fldmiddlename, fldphone, fldemail,
    fldbirthday, fldphoto, flddriverlicensenumber, flddriverlicenseissuedate,
    flddriverlicenseexpirydate, flddriverlicensephoto, flddriverlicensescan,
    fldpassportseries, fldpassportnumber, fldpassportissuedate, fldpassportphoto,
    fldpassportscan, fldaddress, fldexperience, fldstatus
) VALUES
(
    'd001', 'Ivan', 'Ivanov', 'Ivanovich', '+79001234567', 'ivanov@example.com',
    '1985-05-15', '/uploads/drivers/d001.jpg', '1234567890', '2010-06-01',
    '2025-06-01', '/uploads/drivers/license_d001.jpg', '/uploads/drivers/license_scan_d001.pdf',
    '1234', '567890', '2010-05-20', '/uploads/drivers/passport_d001.jpg',
    '/uploads/drivers/passport_scan_d001.pdf', 'г. Москва, ул. Ленина, д. 1, кв. 10', 10, 'active'
),
(
    'd002', 'Petr', 'Petrov', 'Petrovich', '+79002345678', 'petrov@example.com',
    '1988-08-22', '/uploads/drivers/d002.jpg', '2345678901', '2012-07-15',
    '2027-07-15', '/uploads/drivers/license_d002.jpg', '/uploads/drivers/license_scan_d002.pdf',
    '2345', '678901', '2012-07-01', '/uploads/drivers/passport_d002.jpg',
    '/uploads/drivers/passport_scan_d002.pdf', 'г. Санкт-Петербург, ул. Невский, д. 2, кв. 20', 8, 'active'
),
(
    'd003', 'Sergey', 'Sidorov', 'Sergeevich', '+79003456789', 'sidorov@example.com',
    '1990-03-10', '/uploads/drivers/d003.jpg', '3456789012', '2014-09-01',
    '2029-09-01', '/uploads/drivers/license_d003.jpg', '/uploads/drivers/license_scan_d003.pdf',
    '3456', '789012', '2014-08-20', '/uploads/drivers/passport_d003.jpg',
    '/uploads/drivers/passport_scan_d003.pdf', 'г. Новосибирск, ул. Красный проспект, д. 3, кв. 30', 6, 'active'
),
(
    'd004', 'Alexander', 'Kozlov', 'Alexandrovich', '+79004567890', 'kozlov@example.com',
    '1987-11-30', '/uploads/drivers/d004.jpg', '4567890123', '2011-04-10',
    '2026-04-10', '/uploads/drivers/license_d004.jpg', '/uploads/drivers/license_scan_d004.pdf',
    '4567', '890123', '2011-03-25', '/uploads/drivers/passport_d004.jpg',
    '/uploads/drivers/passport_scan_d004.pdf', 'г. Екатеринбург, ул. Мира, д. 4, кв. 40', 9, 'active'
),
(
    'd005', 'Dmitry', 'Morozov', 'Dmitrievich', '+79005678901', 'morozov@example.com',
    '1992-07-18', '/uploads/drivers/d005.jpg', '5678901234', '2016-11-20',
    '2031-11-20', '/uploads/drivers/license_d005.jpg', '/uploads/drivers/license_scan_d005.pdf',
    '5678', '901234', '2016-11-05', '/uploads/drivers/passport_d005.jpg',
    '/uploads/drivers/passport_scan_d005.pdf', 'г. Казань, ул. Баумана, д. 5, кв. 50', 4, 'active'
),
(
    'd006', 'Nikolay', 'Volkov', 'Nikolaevich', '+79006789012', 'volkov@example.com',
    '1989-02-25', '/uploads/drivers/d006.jpg', '6789012345', '2013-08-15',
    '2028-08-15', '/uploads/drivers/license_d006.jpg', '/uploads/drivers/license_scan_d006.pdf',
    '6789', '012345', '2013-08-01', '/uploads/drivers/passport_d006.jpg',
    '/uploads/drivers/passport_scan_d006.pdf', 'г. Нижний Новгород, ул. Горького, д. 6, кв. 60', 7, 'active'
),
(
    'd007', 'Andrey', 'Sokolov', 'Andreevich', '+79007890123', 'sokolov@example.com',
    '1991-09-12', '/uploads/drivers/d007.jpg', '7890123456', '2015-12-01',
    '2030-12-01', '/uploads/drivers/license_d007.jpg', '/uploads/drivers/license_scan_d007.pdf',
    '7890', '123456', '2015-11-15', '/uploads/drivers/passport_d007.jpg',
    '/uploads/drivers/passport_scan_d007.pdf', 'г. Челябинск, ул. Кирова, д. 7, кв. 70', 5, 'active'
),
(
    'd008', 'Maxim', 'Lebedev', 'Maximovich', '+79008901234', 'lebedev@example.com',
    '1986-04-08', '/uploads/drivers/d008.jpg', '8901234567', '2010-10-20',
    '2025-10-20', '/uploads/drivers/license_d008.jpg', '/uploads/drivers/license_scan_d008.pdf',
    '8901', '234567', '2010-10-05', '/uploads/drivers/passport_d008.jpg',
    '/uploads/drivers/passport_scan_d008.pdf', 'г. Омск, ул. Ленина, д. 8, кв. 80', 11, 'active'
),
(
    'd009', 'Alexey', 'Kuznetsov', 'Alexeyevich', '+79009012345', 'kuznetsov@example.com',
    '1993-12-28', '/uploads/drivers/d009.jpg', '9012345678', '2017-03-10',
    '2032-03-10', '/uploads/drivers/license_d009.jpg', '/uploads/drivers/license_scan_d009.pdf',
    '9012', '345678', '2017-02-25', '/uploads/drivers/passport_d009.jpg',
    '/uploads/drivers/passport_scan_d009.pdf', 'г. Самара, ул. Гагарина, д. 9, кв. 90', 3, 'active'
),
(
    'd010', 'Igor', 'Popov', 'Igorovich', '+79010123456', 'popov@example.com',
    '1984-06-05', '/uploads/drivers/d010.jpg', '0123456789', '2009-05-15',
    '2024-05-15', '/uploads/drivers/license_d010.jpg', '/uploads/drivers/license_scan_d010.pdf',
    '0123', '456789', '2009-05-01', '/uploads/drivers/passport_d010.jpg',
    '/uploads/drivers/passport_scan_d010.pdf', 'г. Ростов-на-Дону, ул. Пушкинская, д. 10, кв. 100', 12, 'active'
);
