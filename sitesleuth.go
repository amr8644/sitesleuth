package main

func Start() {

	//tar := "https://dhokla.net/"
	tar := "https://www.shielder.it"
	//ReadFile()
	ScrapeURL(tar)
	CheckHTML(data.HTML)
	for _, v := range data.Script {
		CheckScripts(v)
	}
    
	//CheckHeaders()
	//PrintData()
}
