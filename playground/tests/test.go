package main

import (
	"fmt"
	"strconv"
)

func CreateLink(idx int) (string, string) {
	/*
		This function creates a link for the json using the page id.
	*/
	nmr := strconv.Itoa(idx) // int to string

	path := fmt.Sprintf("%s.json", nmr)
	link_base_1 := "https://www.webmotors.com.br/api/search/car?url=https://www.webmotors.com.br/carros-usados%%2Festoque%%3Finst%%3Dheader%%3Awebmotors%%3Aheader-deslogado%%3A%%3Acarros-usados-ou-seminovos&actualPage="
	link_base_2 := "&displayPerPage=23&order=1&showMenu=true&showCount=true&showBreadCrumb=true&testAB=false&returnUrl=false"

	link := fmt.Sprintf(link_base_1 + nmr + link_base_2)

	return link, path
}

func main() {
	link, path := CreateLink(9999999)
	fmt.Println("\n", link)
	fmt.Println("\n", path)
}
