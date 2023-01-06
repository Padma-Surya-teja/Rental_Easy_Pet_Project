package client

import rental "rental_easy.in/m/pkg/rentalmgmt"

var User1 = rental.User{Name: "Sai Chander", Email: "saichander@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User2 = rental.User{Name: "Rithish Reddy", Email: "ritish.reddy@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User3 = rental.User{Name: "Teja Puvvula", Email: "tejapuvvula@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User4 = rental.User{Name: "Padma Surya Teja", Email: "padmasuryateja@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User5 = rental.User{Name: "Naga Narhsimha", Email: "naga@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User6 = rental.User{Name: "Goutham Pilla", Email: "gp@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User7 = rental.User{Name: "Quizra", Email: "quizra@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User8 = rental.User{Name: "Sandeep Pyata", Email: "sandeep@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User9 = rental.User{Name: "Durga Prakash Madishetty", Email: "dp@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User10 = rental.User{Name: "Mani", Email: "mani@beautifulcode.in", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}
var User11 = rental.User{Name: "Padma surya teja", Email: "19bd1a057d@gmail.com", PhoneNumber: "9999999991", Address: "H-No 1 Near Radhika Theatre Road no 5 Kapra, Secundrabad", District: "Medchal", PostalCode: "500029", Country: "India"}

var Users = []rental.User{User1, User2, User3, User4, User5, User6, User7, User8, User9, User10}

var Item1 = rental.Item{Name: "Asus ZEnBook 14 OLED", Description: "User this Asus zenbook to watch movies in high resolution and Gaming", Category: "Laptops", AvailableFrom: "01-01-2023", AvailableTo: "31-01-2023", AmountPerDay: 1000, UserId: 1}
var Item2 = rental.Item{Name: "Iqoo Neo 6", Description: "Has Snapdragon 870 processor and ois camera", Category: "MobilePhones", AvailableFrom: "01-01-2023", AvailableTo: "31-03-2023", AmountPerDay: 300, UserId: 3}
var Item3 = rental.Item{Name: "Canon EOS 200D II 24.1MP Digital SLR Camera + EF-S 18-55mm f4 is STM Lens (Black)", Description: "Compatible Mountings: Canon Ef-S; Photo Sensor Technology: Size:[Unit:Frames Per Second, Value:Aps-C ], Technology:[Value:Cmos ]; Hardware Interface: Secure Digital", Category: "Cameras", AvailableFrom: "10-01-2023", AvailableTo: "15-01-2023", AmountPerDay: 500, UserId: 6}
var Item4 = rental.Item{Name: "Canon EF50MM F/1.8 STM Lens for Canon DSLR Cameras", Description: "Metal mount with a new exterior design used.The lens is fully compatible with all full range of Canon DSLR Camera.", Category: "Cameras", AvailableFrom: "15-01-2023", AvailableTo: "27-02-2023", AmountPerDay: 300, UserId: 8}
var Item5 = rental.Item{Name: "Canon XA11 Professional Camcorder, Optical, Black", Description: "Full hd 1920x1080 recording capabilities, 3.0-inch lcd capacitive touchscreen with tilt able electronic viewfinder, Genuine canon 20x high definition optical zoom lens", Category: "Cameras", AvailableFrom: "01-01-2023", AvailableTo: "31-01-2023", AmountPerDay: 1000, UserId: 1}
var Item6 = rental.Item{Name: "Pulsor 150", Description: "Gives Mileage of 40kmph", Category: "Bikes", AvailableFrom: "10-01-2023", AvailableTo: "31-01-2023", AmountPerDay: 800, UserId: 8}

var Items = []rental.Item{Item1, Item2, Item3, Item4, Item5, Item6}
