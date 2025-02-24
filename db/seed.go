package db

func SeedDB() {
	plants := []Plant{
		// Takes account leap years
		// January (Days 1-31)
		{Name: "Onion", StartDay: 10, EndDay: 40, Type: "Vegetable"},
		{Name: "Spinach", StartDay: 20, EndDay: 50, Type: "Vegetable"},

		// February (Days 32-59)
		{Name: "Peas", StartDay: 32, EndDay: 71, Type: "Vegetable"},
		{Name: "Kale", StartDay: 40, EndDay: 81, Type: "Vegetable"},
		{Name: "Broccoli", StartDay: 45, EndDay: 91, Type: "Vegetable"},
		{Name: "Asparagus", StartDay: 50, EndDay: 121, Type: "Vegetable"},

		// March (Days 61-90)
		{Name: "Carrots", StartDay: 60, EndDay: 121, Type: "Vegetable"},
		{Name: "Strawberries", StartDay: 60, EndDay: 151, Type: "Fruit"},
		{Name: "Lettuce", StartDay: 65, EndDay: 131, Type: "Vegetable"},
		{Name: "Beets", StartDay: 70, EndDay: 141, Type: "Vegetable"},
		{Name: "Tomatoes", StartDay: 80, EndDay: 151, Type: "Vegetable"},
		{Name: "Chili Peppers", StartDay: 85, EndDay: 171, Type: "Vegetable"},
		{Name: "Bell Peppers", StartDay: 90, EndDay: 161, Type: "Vegetable"},
		{Name: "Eggplant", StartDay: 90, EndDay: 171, Type: "Vegetable"},

		// April (Days 92-120)
		{Name: "Cabbage", StartDay: 100, EndDay: 171, Type: "Vegetable"},
		{Name: "Basil", StartDay: 100, EndDay: 201, Type: "Herb"},
		{Name: "Corn", StartDay: 110, EndDay: 181, Type: "Vegetable"},
		{Name: "Dill", StartDay: 110, EndDay: 211, Type: "Herb"},
		{Name: "Zucchini", StartDay: 120, EndDay: 191, Type: "Vegetable"},
		{Name: "Parsley", StartDay: 120, EndDay: 221, Type: "Herb"},

		// May (Days 122-151)
		{Name: "Cucumbers", StartDay: 130, EndDay: 201, Type: "Vegetable"},
		{Name: "Cilantro", StartDay: 130, EndDay: 231, Type: "Herb"},
		{Name: "Pumpkins", StartDay: 140, EndDay: 221, Type: "Vegetable"},
		{Name: "Green Beans", StartDay: 140, EndDay: 251, Type: "Vegetable"},
		{Name: "Watermelon", StartDay: 150, EndDay: 231, Type: "Fruit"},
		{Name: "Chives", StartDay: 150, EndDay: 261, Type: "Herb"},

		// June (Days 153-181)
		{Name: "Sunflowers", StartDay: 160, EndDay: 241, Type: "Flower"},
		{Name: "Leeks", StartDay: 160, EndDay: 271, Type: "Vegetable"},
		{Name: "Okra", StartDay: 170, EndDay: 261, Type: "Vegetable"},
		{Name: "Thyme", StartDay: 170, EndDay: 281, Type: "Herb"},
		{Name: "Sweet Potatoes", StartDay: 180, EndDay: 271, Type: "Vegetable"},
		{Name: "Rosemary", StartDay: 180, EndDay: 291, Type: "Herb"},

		// July (Days 183-212)
		{Name: "Radish", StartDay: 190, EndDay: 281, Type: "Vegetable"},
		{Name: "Lavender", StartDay: 190, EndDay: 301, Type: "Herb"},
		{Name: "Turnips", StartDay: 200, EndDay: 291, Type: "Vegetable"},
		{Name: "Sage", StartDay: 200, EndDay: 310, Type: "Herb"},
		{Name: "Bok Choy", StartDay: 210, EndDay: 301, Type: "Vegetable"},
		{Name: "Oregano", StartDay: 210, EndDay: 321, Type: "Herb"},

		// August (Days 214-243)
		{Name: "Brussels Sprouts", StartDay: 220, EndDay: 311, Type: "Vegetable"},
		{Name: "Mint", StartDay: 220, EndDay: 331, Type: "Herb"},
		{Name: "Mustard Greens", StartDay: 230, EndDay: 321, Type: "Vegetable"},
		{Name: "Fennel", StartDay: 230, EndDay: 341, Type: "Herb"},
		{Name: "Garlic", StartDay: 240, EndDay: 311, Type: "Vegetable"},
		{Name: "Rhubarb", StartDay: 240, EndDay: 351, Type: "Vegetable"},

		// September (Days 245-273)
		{Name: "Rye", StartDay: 250, EndDay: 271, Type: "Vegetable"},
		{Name: "Swiss Chard", StartDay: 250, EndDay: 361, Type: "Vegetable"},
		{Name: "Collard Greens", StartDay: 260, EndDay: 271, Type: "Vegetable"},
		{Name: "Horseradish", StartDay: 270, EndDay: 301, Type: "Vegetable"},

		// October (Days 275-304)
		{Name: "Endive", StartDay: 280, EndDay: 301, Type: "Vegetable"},
		{Name: "Arugula", StartDay: 290, EndDay: 301, Type: "Vegetable"},
		{Name: "Celery", StartDay: 300, EndDay: 331, Type: "Vegetable"},
		{Name: "Winter Lettuce", StartDay: 300, EndDay: 331, Type: "Vegetable"},

		// November (Days 306-334)
		{Name: "Cress", StartDay: 310, EndDay: 331, Type: "Herb"},
		{Name: "Herbs (Indoor)", StartDay: 310, EndDay: 61, Type: "Herb"},
		{Name: "Mushrooms", StartDay: 320, EndDay: 361, Type: "Vegetable"},

		// December (Days 336-365)
		{Name: "Fava Beans", StartDay: 335, EndDay: 31, Type: "Vegetable"},
		{Name: "Winter Peas", StartDay: 335, EndDay: 31, Type: "Vegetable"},
	}

	if err := DB.Save(&plants).Error; err != nil {
		panic("Failed to seed database: " + err.Error())
	}
}
