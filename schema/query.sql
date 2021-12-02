-- name: GetAllContacts :many
SELECT contacts.*,address,zip,city,country,mobile,email  FROM contacts LEFT JOIN contacts_address ON  contacts.contact_id = contacts_address.contact_id  AND contacts_address.address_type='default' LIMIT @result_limit::int  OFFSET @result_offset::int;

-- name: GetAllContactsNew :many
SELECT contacts.*,address,zip,city,country,mobile,email  FROM contacts LEFT JOIN contacts_address ON  contacts.contact_id = contacts_address.contact_id  AND contacts_address.address_type='default';


-- name: GetContactCount :one
SELECT count(*) FROM contacts;


-- name: GetAllContactstest :many
SELECT * FROM contacts;

-- name: CreateContact :one
INSERT INTO contacts(
	contact_group_id,
	first_name,
	last_name,
	business_name,
	job_title,
	location_group_id,
	image,
	description
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING  contact_id;

-- name: UpdateContact :one
UPDATE contacts SET  
    first_name = $2, 
    last_name = $3,
    business_name = $4,
    job_title = $5,
    location_group_id = $6,
	image = $7,
    description = $8
  WHERE contact_id = $1 RETURNING contact_id; 



-- name: GetContactById :one
SELECT * FROM contacts WHERE contact_id = $1;

-- name: DeleteContactById :exec
DELETE FROM contacts WHERE contact_id = $1;







-- name: GetAllContactAddress :many
SELECT * FROM contacts_address;

-- name: GetContactAddressById :one
SELECT * FROM contacts_address WHERE contacts_address_id = $1;

-- name: GetContactAddressByContact :many
SELECT * FROM contacts_address WHERE contact_id = $1;

-- name: GetContactAddressId :one
SELECT contacts_address_id FROM contacts_address WHERE contact_id = $1 and address_type = $2 limit 1;


-- name: CreateContactAddress :one
INSERT INTO contacts_address(
	contact_id,
	address_type,
	address,
	zip,
	city ,
	country,
	mobile ,
	telephone,
	email,
	website
) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING  contacts_address_id;

-- name: UpdateContactAddress :one
UPDATE contacts_address 
SET  
	website = $3,
	address = $4,
	zip = $5,
	city = $6,
	country = $7,
	mobile = $8 ,
	telephone =$9,
	email = $10
	
	WHERE contacts_address_id = $1 and contact_id = $2 RETURNING contacts_address_id;

-- name: UpdateContactType :one
UPDATE contact_type 
SET  is_active = $2 , is_location_group = $3  WHERE contact_type_id = $1 RETURNING contact_type_id ; 

-- name: GetAllContactType :many
SELECT contact_type_id, contact_type_name, is_active, is_location_group FROM contact_type order by contact_type_name asc;

-- name: CreateContactTypeSubGroup :one
INSERT INTO contact_group(
	contact_group_name,
	contact_type_id
) VALUES ($1,$2) RETURNING  contact_group_id;

-- name: GetAllContactTypeSubGroup :many
SELECT * FROM contact_group;

-- name: GetContactTypeSubGroupById :one
SELECT * FROM contact_group WHERE contact_group_id = $1;

-- name: UpdateContactTypeSubGroup :one
UPDATE contact_group 
SET  contact_group_name = $2 , contact_type_id = $3,is_active = $4   WHERE contact_group_id = $1 RETURNING contact_group_id; 
