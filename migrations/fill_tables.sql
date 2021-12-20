-- Insert initial data in status table
INSERT INTO public.status(name, description)
VALUES 
('estimate', 'Estimate'),
('issued', 'Issued'),
('in_progress', 'In Progress'),
('fulfilled', 'Fulfilled'),
('closed_short', 'Closed Short'),
('void', 'Void'),
('expired', 'Expired');