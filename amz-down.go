package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var files = []string{
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/AdditionalProductInformation.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Amazon.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/amzn-base.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/amzn-envelope.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/amzn-header.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Arts.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/AttributeGroups.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/AutoAccessory.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Baby.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Beauty.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Books.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/CameraPhoto.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/CatPIL.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/CE.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ClothingAccessories.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Coins.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Collectibles.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Computers.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Customer.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/CustomerAddress.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/EducationSupplies.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/EntertainmentCollectibles.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/EUCompliance.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/FBA.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/FoodAndBeverages.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/FoodServiceAndJanSan.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/FulfillmentCenter.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/FulfillmentOrderCancellationRequest.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/FulfillmentOrderRequest.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Furniture.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/GiftCards.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Gourmet.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Health.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Home.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/HomeImprovement.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Image.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Industrial.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Inventory.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Item.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Jewelry.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/LabSupplies.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/LargeAppliances.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Lighting.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/LightMotor.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Listings.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ListingSummary.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Loyalty.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Luggage.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/LuxuryBeauty.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/MaterialHandling.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/MechanicalFasteners.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Miscellaneous.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Miscellaneous.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Motorcycles.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/MultiChannelOrderReport.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Music.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/MusicalInstruments.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/NavigationReport.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Offer.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Office.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/OrderAcknowledgement.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/OrderAdjustment.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/OrderFulfillment.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/OrderNotificationReport.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/OrderReport.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Outdoors.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Override.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/PaymentMethod.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/PetSupplies.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/PowerTransmission.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Price.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ProcessingReport.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Product.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ProductAttributes.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ProductClothing.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ProductImage.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/ProfessionalHealthCare.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/RawMaterials.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Relationship.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ReverseFeed.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/SettlementReport.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/Shoes.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Sports.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_9/SportsMemorabilia.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Store.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/SWVG.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_1_7/TestOrderRequest.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ThreeDPrinting.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/TiresAndWheels.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Tools.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Toys.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/ToysBaby.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/TypeDefinitions.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Video.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/WebstoreItem.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/WineAndAlcohol.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/rainier/help/xsd/release_4_1/Wireless.xsd",
	"https://www.w3.org/TR/xmlenc-core/xenc-schema.xsd",
	"https://www.w3.org/TR/2002/REC-xmldsig-core-20020212/xmldsig-core-schema.xsd",
	"https://images-na.ssl-images-amazon.com/images/G/01/mwsportal/doc/en_US/Reports/XSDs/BrowseTreeReport.xsd",
}

var files_name = []string{
	"AdditionalProductInformation.xsd",
	"Amazon.xsd",
	"amzn-base.xsd",
	"amzn-envelope.xsd",
	"amzn-header.xsd",
	"Arts.xsd",
	"AttributeGroups.xsd",
	"AutoAccessory.xsd",
	"Baby.xsd",
	"Beauty.xsd",
	"Books.xsd",
	"CameraPhoto.xsd",
	"CatPIL.xsd",
	"CE.xsd",
	"ClothingAccessories.xsd",
	"Coins.xsd",
	"Collectibles.xsd",
	"Computers.xsd",
	"Customer.xsd",
	"CustomerAddress.xsd",
	"EducationSupplies.xsd",
	"EntertainmentCollectibles.xsd",
	"EUCompliance.xsd",
	"FBA.xsd",
	"FoodAndBeverages.xsd",
	"FoodServiceAndJanSan.xsd",
	"FulfillmentCenter.xsd",
	"FulfillmentOrderCancellationRequest.xsd",
	"FulfillmentOrderRequest.xsd",
	"Furniture.xsd",
	"GiftCards.xsd",
	"Gourmet.xsd",
	"Health.xsd",
	"Home.xsd",
	"HomeImprovement.xsd",
	"Image.xsd",
	"Industrial.xsd",
	"Inventory.xsd",
	"Item.xsd",
	"Jewelry.xsd",
	"LabSupplies.xsd",
	"LargeAppliances.xsd",
	"Lighting.xsd",
	"LightMotor.xsd",
	"Listings.xsd",
	"ListingSummary.xsd",
	"Loyalty.xsd",
	"Luggage.xsd",
	"LuxuryBeauty.xsd",
	"MaterialHandling.xsd",
	"MechanicalFasteners.xsd",
	"Miscellaneous.xsd",
	"Miscellaneous.xsd",
	"Motorcycles.xsd",
	"MultiChannelOrderReport.xsd",
	"Music.xsd",
	"MusicalInstruments.xsd",
	"NavigationReport.xsd",
	"Offer.xsd",
	"Office.xsd",
	"OrderAcknowledgement.xsd",
	"OrderAdjustment.xsd",
	"OrderFulfillment.xsd",
	"OrderNotificationReport.xsd",
	"OrderReport.xsd",
	"Outdoors.xsd",
	"Override.xsd",
	"PaymentMethod.xsd",
	"PetSupplies.xsd",
	"PowerTransmission.xsd",
	"Price.xsd",
	"ProcessingReport.xsd",
	"Product.xsd",
	"ProductAttributes.xsd",
	"ProductClothing.xsd",
	"ProductImage.xsd",
	"ProfessionalHealthCare.xsd",
	"RawMaterials.xsd",
	"Relationship.xsd",
	"ReverseFeed.xsd",
	"SettlementReport.xsd",
	"Shoes.xsd",
	"Sports.xsd",
	"SportsMemorabilia.xsd",
	"Store.xsd",
	"SWVG.xsd",
	"TestOrderRequest.xsd",
	"ThreeDPrinting.xsd",
	"TiresAndWheels.xsd",
	"Tools.xsd",
	"Toys.xsd",
	"ToysBaby.xsd",
	"TypeDefinitions.xsd",
	"Video.xsd",
	"WebstoreItem.xsd",
	"WineAndAlcohol.xsd",
	"Wireless.xsd",
	"xenc-schema.xsd",
	"xmldsig-core-schema.xsd",
	"BrowseTreeReport.xsd",
}

func main() {
	err := os.MkdirAll("./amz-xsd", os.ModePerm)
	if err != nil {
		panic(err)
	}
	for i, url := range files {
		err := DownloadFile("./amz-xsd/"+files_name[i], url)
		if err != nil {
			panic(err)
		} else {
			fmt.Println(url)
		}
	}
}

func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
