// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"reflect"
	"testing"
)

func TestFrontpage_type(t *testing.T) {
	var a Frontpage
	a.Items = append(a.Items, frontpageItemType{})
	_ = a
}

func TestFrontpage_parseDML(t *testing.T) {
	var dml FrontpageDML
	if err := dml.ParseJson([]byte(testJsonDataFrontpage)); err != nil {
		t.Fatal(err)
	}
	var page Frontpage
	if err := page.ParseDML(&dml); err != nil {
		t.Fatal(err)
	}

	items := page.Items
	page.Items = nil

	if !reflect.DeepEqual(testGoldenFrontpage, page) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", testGoldenFrontpage, page)
	}
	for i := 0; i < len(items) && i < len(testGoldenFrontpageItems); i++ {
		items[i].Description = "" // too large, ingore
		items[i].TextSummary = "" // too large, ingore
		if !reflect.DeepEqual(testGoldenFrontpageItems[i], frontpageItemType(items[i])) {
			t.Fatalf("not equal, expect = \n%q, got = \n%q", testGoldenFrontpageItems[i], items[i])
		}
	}
}

var testGoldenFrontpage = Frontpage{
	Id:        0,
	Title:     "Breaking News and Opinion on The Huffington Post",
	SourceURL: "http://www.huffingtonpost.com",
	Icon:      "http://www.huffingtonpost.com:80/favicon.ico",
	NumItems:  51,
}
var testGoldenFrontpageItems = []frontpageItemType{
	{
		Id:          180194704,
		Title:       "The Austerity Trap and the Jobs Deficit",
		Description: "", // too large, ingore
		XRoot:       "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[1]",
		PubDate:     "Thu, 17 May 2012 20:06:26 GMT",
		Link:        "http://www.huffingtonpost.com:80/robert-l-borosage/the-austerity-trap-and-th_b_1524662.html",
		Type:        "STORY",
		Img:         "",
		TextSummary: "", // too large, ingore
		Sp:          0.000,
		Sr:          4.000,
		Fresh:       1.000,
	},
	{
		Id:          -91871119,
		Title:       "Bringing Water And Sanitation To Ethiopian Children",
		Description: "", // too large, ingore
		XRoot:       "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[2]",
		PubDate:     "Thu, 17 May 2012 20:06:26 GMT",
		Link:        "http://at.atwola.com:80/adlink/3.0/5113.1/2298499/1/16/AdId=2520773;BnId=1;link=http://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.htmlhttp://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.html",
		Type:        "STORY",
		Img:         "",
		TextSummary: "", // too large, ingore
		Sp:          0.101,
		Sr:          4.000,
		Fresh:       1.000,
	},
}

const testJsonDataFrontpage = `
{
  "id": "0",
  "tagName": "dml",
  "childNodes": [
    {
      "tagName": "info",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Breaking News and Opinion on The Huffington Post"
          ]
        },
        {
          "tagName": "sourceType",
          "childNodes": [
            "html"
          ]
        },
        {
          "tagName": "sourceURL",
          "childNodes": [
            "http://www.huffingtonpost.com"
          ]
        },
        {
          "tagName": "icon",
          "childNodes": [
            "http://www.huffingtonpost.com:80/favicon.ico"
          ]
        },
        {
          "tagName": "numItems",
          "childNodes": [
            "51"
          ]
        },
        {
          "tagName": "numSpamItems",
          "childNodes": [
            "0"
          ]
        }
      ]
    },
    {
      "id": "180194704",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "The Austerity Trap and the Jobs Deficit"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/robert-l-borosage/the-austerity-trap-and-th_b_1524662.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Before dealing with the \"prairie fire\" that threatens the nation, Romney and Republicans want to add fuel to the flames. Their first priority is spending more money on the military and collecting less money from the rich and the corporations."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1524662,\"mpid\":2}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/robert-l-borosage\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Robert L. Borosage\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/robert-l-borosage/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/robert-l-borosage/the-austerity-trap-and-th_b_1524662.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">The Austerity Trap and the Jobs Deficit</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/robert-l-borosage\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Robert L. Borosage</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">President, Institute for America's Future</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Before dealing with the \"prairie fire\" that threatens the nation, Romney and Republicans want to add fuel to the flames. Their first priority is spending more money on the military and collecting less money from the rich and the corporations.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/robert-l-borosage/the-austerity-trap-and-th_b_1524662.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/robert-l-borosage/the-austerity-trap-and-th_b_1524662.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(34)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "34",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[1]"
    },
    {
      "id": "-91871119",
      "sp": "0.101",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Bringing Water And Sanitation To Ethiopian Children"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://at.atwola.com:80/adlink/3.0/5113.1/2298499/1/16/AdId=2520773;BnId=1;link=http://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.htmlhttp://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Johnson & Johnson, New Jersey Council for the Humanities, Creative Writing, Parenting, ACCP, PQMD, Head Start"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div style=\"color: #000; text-align: left; font: 12px Georgia;\">\n\t<div>\n\t   \n\t\t\n<div style=\"color: #000; -khtml-border-radius: 5px; -moz-border-radius: 5px; -webkit-border-radius: 5px; border: 1px solid #949494; border-radius: 5px; margin: 12px 0 2px 0; overflow: hidden; padding-top: 8px; width: 300px; text-align: left; font: 12px Georgia;\">\n\t\t<div style=\"-khtml-border-radius: 5px; -moz-border-radius: 5px; -webkit-border-radius: 5px; border-radius: 5px; font: bold 10px/15px Arial,Helvetica,sans-serif; padding: 2px 8px; position: absolute; text-decoration: none; text-transform: uppercase; margin-left: 8px; background-color: #f61784 !important;\">\n                                Global Motherhood            \t</div>\t<ul style=\"overflow: hidden; width: 300px;\">\n\t\t\t<li style=\"color: #000; list-style: none; margin: 4px 0; padding: 4px 8px 8px 8px; text-align: left; width: 284px; border-top: none; font: 12px Georgia; display: inline-block;\">\n\t\t\t\t\t\t\t\t\t\t<div style=\"font-family: Georgia,Century,Times,serif !important; font-size: 13px; margin: 0; width: 290px; list-style-type: disc; border-bottom: none; padding-bottom: 0;\">\n\t\t\t\t\t<div style=\"color: #000; list-style: none; text-align: left; clear: both; font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; font: 12px Georgia;\"></div>\t\t\t\t\t\n\t\t\t\t\t<div style=\"color: #000; list-style: none; text-align: left; float: left; width: 285px; font: 12px Georgia; margin-top: 5px;\">\n\t\t\t\t\t\t<div style=\"float: right; font-family: Georgia,Century,Times,serif !important; font-size: 13px; width: 50px; list-style-type: disc; padding: 0pt 5px;\">\t\t\t\t\t\t\t\t\n\t\t\t\t\t\t\t<img alt=\"Conrad Person\" src=\"http://www.huffingtonpost.com/contributors/conrad-person/headshot.jpg\" style=\"color: #000; list-style: none; text-align: left; font: 12px Georgia; width: 45;\" width=\"45\"></img>\n\t\t\t\t\t\t</div>\n\t\t\t\t\t\t<span style=\"font-family: Georgia,Century,Times,serif !important; font-size: 13px; list-style-type: disc; width: 225px;\"><a href=\"http://at.atwola.com/adlink/3.0/5113.1/2298499/1/16/AdId=2520773;BnId=1;link=http://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.htmlhttp://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.html\" style=\"list-style: none; text-align: left; outline: medium none; text-decoration: none; color: #111111; font-size: 14.03px; font-weight: bold; font: 12px Georgia;\">Bringing Water And Sanitation To Ethiopian Children</a></span>\n\t\t\t\t\t\t<p style=\"font-family: Georgia,Century,Times,serif !important; font-size: 13px; margin: 5px 0 0 0 !important; list-style-type: disc; width: 225px;\"><a href=\"http://www.huffingtonpost.com/conrad-person\" style=\"list-style: none; text-align: left; outline: medium none; text-decoration: none; font-weight: bold; line-height: 18px; color: #ED0978 !important; font: 12px Georgia;\">Conrad Person</a>, 05.17.2012</p>\n\t\t\t\t\t\t<p style=\"font-family: Georgia,Century,Times,serif !important; font-size: 13px; font-style: italic !important; padding-bottom: 5px; list-style-type: disc; width: 235px;\">Johnson &amp; Johnson, New Jersey Council for the Humanities, Creative Writing, Parenting, ACCP, PQMD, Head Start</p>\n\t\t\t\t\t</div>\n\t\t\t\t\t<div style=\"color: #000; list-style: none; text-align: left; font: 12px Georgia; clear: both;\"></div>\n\t\t\t\t\t<p style=\"color: #000; list-style: none; text-align: left; font: 12px Georgia; margin: 2px 0; width: 285px;\">Whether you live in the United States or in Ethiopia, sanitation and proper hygiene is an issue that effects you, but in Ethiopia, the lack of convenient water sources can make the difference between life and death.</p>\n\t\t\t\t\t\t\t\t\t\t<div style=\"color: #000; list-style: none; text-align: left; padding-top: 8px; font: 12px Georgia;\">\n\t\t\t\t\t\t<p style=\"font-family: Georgia,Century,Times,serif !important; font-size: 13px; list-style-type: disc;\">\n\t\t\t\t\t\t\t<a href=\"http://at.atwola.com/adlink/3.0/5113.1/2298499/1/16/AdId=2520773;BnId=1;link=http://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.htmlhttp://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.html\" style=\"list-style: none; text-align: left; outline: medium none; text-decoration: none; font-family: Arial,Helvetica,sans-serif; font-size: 99%; color: #000000 !important; font-weight: bold; font: 12px Georgia;\">Read Post</a>\n\t\t\t\t\t\t\t| <a href=\"http://at.atwola.com/adlink/3.0/5113.1/2298499/1/16/AdId=2520773;BnId=1;link=http://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.htmlhttp://www.huffingtonpost.com/conrad-person/caring-for-our-future-bri_b_1521005.html#comments\" style=\"list-style: none; text-align: left; outline: medium none; text-decoration: none; color: #444444; font-family: Arial,Helvetica,sans-serif; font-size: 99%; font: 12px Georgia;\">Comments <span style=\"font-family: Arial,Helvetica,sans-serif !important; list-style-type: disc;\"></span></a>\n\t\t\t\t\t\t</p>\n\t\t\t\t\t</div>\n\t\t\t\t\t\t\t\t\t</div>\n\t\t\t\t\t</li>\n\t\t</ul>\n\t\t<div style=\"clear: both;\"></div>\n</div>\n\n\n\t</div>\n</div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[2]"
    },
    {
      "id": "1041056161",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Tech Apps: A Few of My Favorite Things"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/marlo-thomas/tech-apps-a-few-of-my-fav_b_1500088.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I'm all plugged in at this point. And now I'm constantly discovering new apps, gadgets and gizmos that I just love."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1500088,\"mpid\":3}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/marlo-thomas\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Marlo Thomas\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/marlo-thomas/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/women/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #BA72BA !important; text-align: left; font: 12px Georgia;\">Women</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/marlo-thomas/tech-apps-a-few-of-my-fav_b_1500088.html\" style=\"outline: none; text-decoration: none; color: #BA72BA !important;\">Tech Apps: A Few of My Favorite Things</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/marlo-thomas\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Marlo Thomas</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Award-winning actress, author, and activist</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I'm all plugged in at this point. And now I'm constantly discovering new apps, gadgets and gizmos that I just love.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/marlo-thomas/tech-apps-a-few-of-my-fav_b_1500088.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/marlo-thomas/tech-apps-a-few-of-my-fav_b_1500088.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(13)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "13",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[3]"
    },
    {
      "id": "-607289475",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Preying on the Poor"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/barbara-ehrenreich/us-poverty_b_1523865.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Lenders, including major credit companies as well as payday lenders, have taken over the traditional role of the street-corner loan shark, charging the poor insanely high rates of interest."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1523865,\"mpid\":4}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/barbara-ehrenreich\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Barbara Ehrenreich\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/barbara-ehrenreich/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/barbara-ehrenreich/us-poverty_b_1523865.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">Preying on the Poor</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/barbara-ehrenreich\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Barbara Ehrenreich</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">New York Times bestselling author</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Lenders, including major credit companies as well as payday lenders, have taken over the traditional role of the street-corner loan shark, charging the poor insanely high rates of interest.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/barbara-ehrenreich/us-poverty_b_1523865.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/barbara-ehrenreich/us-poverty_b_1523865.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(300)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "300",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[4]"
    },
    {
      "id": "-1450310588",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Food Revolution Day : Jamie Oliver ... Jamie Oliver, 05/17/12 ... Celebrity Chef and Food Activist"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/jamie-oliver/food-revolution-day_b_1524926.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1524926,\"mpid\":5}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/jamie-oliver\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Jamie Oliver\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/jamie-oliver/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/food/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #5C3545 !important; text-align: left; font: 12px Georgia;\">Food</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/jamie-oliver/food-revolution-day_b_1524926.html\" style=\"outline: none; text-decoration: none; color: #5C3545 !important;\">Food Revolution Day</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/jamie-oliver\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Jamie Oliver</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Celebrity Chef and Food Activist</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">The Food Revolution and Food Revolution Day is about empowering people through education or, frankly, just inspiring people to be more street-wise about food, where it comes from and how it affects their bodies.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/jamie-oliver/food-revolution-day_b_1524926.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/jamie-oliver/food-revolution-day_b_1524926.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "LINK",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[5]"
    },
    {
      "id": "-361167782",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Donna Summer, 12/31/48 - 5/17/12"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/hank-bordowitz/donna-summer-dead_b_1525015.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "More than any other performer of the mid-70s through the 80s, Donna Summer's recordings pushed the envelope of dance music."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1525015,\"mpid\":6}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/hank-bordowitz\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Hank Bordowitz\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/hank-bordowitz/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/fifty/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #7e5bff; text-align: left; font: 12px Georgia;\">Post50</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/hank-bordowitz/donna-summer-dead_b_1525015.html\" style=\"outline: none; text-decoration: none; color: #7e5bff;\">Donna Summer, 12/31/48 - 5/17/12</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/hank-bordowitz\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Hank Bordowitz</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Author</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">More than any other performer of the mid-70s through the 80s, Donna Summer's recordings pushed the envelope of dance music.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/hank-bordowitz/donna-summer-dead_b_1525015.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/hank-bordowitz/donna-summer-dead_b_1525015.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[7]"
    },
    {
      "id": "361755709",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Because They Can"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/james-kwak/national-debt_b_1523771.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Associate professor, University of Connecticut School of Law, and co-author, 'White House Burning'"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1523771,\"mpid\":7}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/james-kwak\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"James Kwak\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/james-kwak/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/business/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #1F0050 !important; text-align: left; font: 12px Georgia;\">Business</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/james-kwak/national-debt_b_1523771.html\" style=\"outline: none; text-decoration: none; color: #1F0050 !important;\">Because They Can</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/james-kwak\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">James Kwak</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Associate professor, University of Connecticut School of Law, and co-author, 'White House Burning'</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">It seems as if the Republicans are trying to turn the national debt back into a major political issue. Despite their abysmal record when it comes to fiscal responsibility, it could still turn out to be smart politics, for a few reasons.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/james-kwak/national-debt_b_1523771.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/james-kwak/national-debt_b_1523771.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(109)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "109",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[8]"
    },
    {
      "id": "1060972331",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "NDAA Authorizes War Against Iran"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/rep-dennis-kucinich/ndaa-authorizes-war-again_b_1524474.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "This week, Congress is considering two pieces of legislation relating to Iran. The first undermines a diplomatic solution with Iran and lowers the bar for war."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1524474,\"mpid\":8}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/rep-dennis-kucinich\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Rep. Dennis Kucinich\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/rep-dennis-kucinich/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/rep-dennis-kucinich/ndaa-authorizes-war-again_b_1524474.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">NDAA Authorizes War Against Iran</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/rep-dennis-kucinich\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Rep. Dennis Kucinich</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">U.S. Representative from Ohio's 10th District</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">This week, Congress is considering two pieces of legislation relating to Iran. The first undermines a diplomatic solution with Iran and lowers the bar for war. The second authorizes a war of choice against Iran and begins military preparations for it.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/rep-dennis-kucinich/ndaa-authorizes-war-again_b_1524474.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/rep-dennis-kucinich/ndaa-authorizes-war-again_b_1524474.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(18)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "18",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[9]"
    },
    {
      "id": "-1534913246",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "5.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "American Idol Recap: The Final Three Perform"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/laura-prudom/american-idol-recap-final-three-perform_b_1522982.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I gave up on being invested in the outcome of an Idol season the year that the phenomenally talented Adam Lambert lost to the pretty-but-bland Kris Allen, so I honestly have no horse in this race. But if I had to bet, I'd bet on Phillip Phillips."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1522982,\"mpid\":9}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/laura-prudom\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Laura Prudom\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/laura-prudom/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/tv/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #760101 !important; text-align: left; font: 12px Georgia;\">TV</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/laura-prudom/american-idol-recap-final-three-perform_b_1522982.html\" style=\"outline: none; text-decoration: none; color: #760101 !important;\"><i style=\"font-style: italic !important; line-height: 16px; text-align: left; font: 12px Georgia; font-weight: bold;\">American Idol</i> Recap: The Final Three Perform</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/laura-prudom\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Laura Prudom</a>, 05/16/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Associate Television Editor, The Huffington Post</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I gave up on being invested in the outcome of an <em style=\"color: #000; font-style: italic !important; text-align: left; font: 12px Georgia;\">Idol</em> season the year that the phenomenally talented Adam Lambert lost to the pretty-but-bland Kris Allen, so I honestly have no horse in this race. But if I had to bet, I'd bet on Phillip Phillips.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/laura-prudom/american-idol-recap-final-three-perform_b_1522982.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/laura-prudom/american-idol-recap-final-three-perform_b_1522982.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(71)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "71",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[10]"
    },
    {
      "id": "-316661148",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Black-on-White Crime and the Reasons for a Media Double-Standard"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/bob-cesca/blackonwhite-crime-and-th_b_1521775.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Yes, there's a double-standard. And until there's full equality and the long slow process of racial healing is completed, the double-standard has to remain."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1521775,\"mpid\":10}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/bob-cesca\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Bob Cesca\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/bob-cesca/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/media/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #DD5200 !important; text-align: left; font: 12px Georgia;\">Media</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/bob-cesca/blackonwhite-crime-and-th_b_1521775.html\" style=\"outline: none; text-decoration: none; color: #DD5200 !important;\">Black-on-White Crime and the Reasons for a Media Double-Standard</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/bob-cesca\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Bob Cesca</a>, 05/16/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Host of The Bob &amp; Chez Show, Media Producer</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Yes, there's a double-standard. And until there's full equality and the long slow process of racial healing is completed, the double-standard <em style=\"color: #000; font-style: italic !important; text-align: left; font: 12px Georgia;\">has to</em> remain.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/bob-cesca/blackonwhite-crime-and-th_b_1521775.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/bob-cesca/blackonwhite-crime-and-th_b_1521775.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(676)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "676",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[11]"
    },
    {
      "id": "-600635503",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "5.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Kiss Me, but Finish Singing That High Note First"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/jennifer-rivera/opera-kissing_b_1524284.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I always prefer to get the kissing out of the way right away so that the awkwardness dissolves by opening night. Plus, it's better to not be surprised by the fact that your tenor is actually an open mouthed wet kisser, or that your soprano is a closed lip face presser."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1524284,\"mpid\":11}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/jennifer-rivera\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Jennifer Rivera\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/jennifer-rivera/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/culture/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #28787f !important; text-align: left; font: 12px Georgia;\">Culture</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/jennifer-rivera/opera-kissing_b_1524284.html\" style=\"outline: none; text-decoration: none; color: #28787f !important;\">Kiss Me, but Finish Singing That High Note First</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/jennifer-rivera\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Jennifer Rivera</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Mezzo soprano</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I always prefer to get the kissing out of the way right away so that the awkwardness dissolves by opening night. Plus, it's better to not be surprised by the fact that your tenor is actually an open mouthed wet kisser, or that your soprano is a closed lip face presser.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/jennifer-rivera/opera-kissing_b_1524284.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/jennifer-rivera/opera-kissing_b_1524284.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[12]"
    },
    {
      "id": "-1604519136",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "5.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "No, President Obama Is Not the 'First Gay President'"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/michelangelo-signorile/no-president-obama-is-not_b_1524125.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Even if you accept Toni Morrison's premise that Bill Clinton was \"the first black president,\" which of course Andrew Sullivan is working off of in his essay in Newsweek, it just falls flat when applied to Obama and gays."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1524125,\"mpid\":12}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/michelangelo-signorile\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Michelangelo Signorile\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/michelangelo-signorile/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/gay-voices/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #ED4A4B !important; text-align: left; font: 12px Georgia;\">Gay Voices</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/michelangelo-signorile/no-president-obama-is-not_b_1524125.html\" style=\"outline: none; text-decoration: none; color: #ED4A4B !important;\">No, President Obama Is Not the 'First Gay President'</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/michelangelo-signorile\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Michelangelo Signorile</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Editor-at-large, HuffPost Gay Voices; SiriusXM radio host</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Even if you accept Toni Morrison's premise that Bill Clinton was \"the first black president,\" which of course Andrew Sullivan is working off of in his essay in <em style=\"color: #000; font-style: italic !important; text-align: left; font: 12px Georgia;\">Newsweek</em>, it just falls flat when applied to Obama and gays.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/michelangelo-signorile/no-president-obama-is-not_b_1524125.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/michelangelo-signorile/no-president-obama-is-not_b_1524125.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(23)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "23",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[13]"
    },
    {
      "id": "-811364529",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "5.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "How to Prevent Indefinite Military Detention in the United States. Really."
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/daphne-eviatar/how-to-prevent-indefinite_b_1525037.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "The National Defense Authorization Act rightly outraged many on both the left and the right last year, and legislators from all over the political map are now responding. But while one of those responses is real -- i."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1525037,\"mpid\":13}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/daphne-eviatar\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Daphne Eviatar\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/daphne-eviatar/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/daphne-eviatar/how-to-prevent-indefinite_b_1525037.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">How to Prevent Indefinite Military Detention in the United States. Really.</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/daphne-eviatar\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Daphne Eviatar</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Senior Counsel, Human Rights First's Law and Security Program</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">The National Defense Authorization Act rightly outraged many on both the left and the right last year, and legislators from all over the political map are now responding.  But while one of those responses is real -- i.e., it would actually fix the problem -- one is not, and by pretending to fix it would only make things worse.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/daphne-eviatar/how-to-prevent-indefinite_b_1525037.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/daphne-eviatar/how-to-prevent-indefinite_b_1525037.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[14]"
    },
    {
      "id": "-2024599554",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Innocent & Executed: It Could Have Been Me"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/franky-carrillo/carlos-deluna-execution-innocent_b_1525328.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I was wrongfully convicted when I was 16 years old and served 20 years in prison before proving my innocence. That mistake took two decades from me, but it took Carlos DeLuna's life."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1525328,\"mpid\":14}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/franky-carrillo\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Franky Carrillo\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/franky-carrillo/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/crime/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #513420 !important; text-align: left; font: 12px Georgia;\">Crime</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/franky-carrillo/carlos-deluna-execution-innocent_b_1525328.html\" style=\"outline: none; text-decoration: none; color: #513420 !important;\">Innocent &amp; Executed: It Could Have Been Me</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/franky-carrillo\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Franky Carrillo</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Exonerated of murder after 20 years in prison</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I was wrongfully convicted when I was 16 years old and served 20 years in prison before proving my innocence. That mistake took two decades from me, but it took Carlos DeLuna's life.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/franky-carrillo/carlos-deluna-execution-innocent_b_1525328.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/franky-carrillo/carlos-deluna-execution-innocent_b_1525328.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[15]"
    },
    {
      "id": "1656626748",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "The History of the Windmill Theatre: The Naked Women Who Defied Hitler"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/judith-walkowitz/windmill-theatre-history_b_1522007.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Despite its sleazy exoticism, London's Soho in the 1930s boasted one local institution that catered to middlebrow suburbia."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1522007,\"mpid\":15}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/judith-walkowitz\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Judith Walkowitz\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/judith-walkowitz/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/books/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #E61405 !important; text-align: left; font: 12px Georgia;\">Books</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html\" style=\"outline: none; text-decoration: none; color: #E61405 !important;\">The History of the Windmill Theatre: The Naked Women Who Defied Hitler</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/judith-walkowitz\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Judith Walkowitz</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Professor of History</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Despite its sleazy exoticism, London's Soho in the 1930s boasted one local institution that catered to middlebrow suburbia.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(114)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "114",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[16]"
    },
    {
      "id": "-969046865",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Talk Nerdy to Me: Mad Science: Giovanni Aldini, Corpse Reanimator"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/15/mad-science-giovanni-aldini_n_1519723.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "2012-05-16-Screenshot20120516at3.52.36PM.pngGiovanni Aldini may not be a household name, but his contributions to science cannot be ignored."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1519723,\"mpid\":16}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/cara-santa-maria\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Cara Santa Maria\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/cara-santa-maria/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/science/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #4e677b !important; text-align: left; font: 12px Georgia;\">Science</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/2012/05/15/mad-science-giovanni-aldini_n_1519723.html\" style=\"outline: none; text-decoration: none; color: #4e677b !important;\"><i style=\"font-style: italic !important; line-height: 16px; text-align: left; font: 12px Georgia; font-weight: bold;\">Talk Nerdy to Me</i>: Mad Science: Giovanni Aldini, Corpse Reanimator</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/cara-santa-maria\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Cara Santa Maria</a>, 05/15/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Science correspondent, Huffington Post; editor, Talk Nerdy to Me</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/2012/05/15/mad-science-giovanni-aldini_n_1519723.html?1337197338\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"2012-05-16-Screenshot20120516at3.52.36PM.png\" height=\"62\" src=\"http://images.huffingtonpost.com/2012-05-16-Screenshot20120516at3.52.36PM.png\" style=\"font-size: 13px; line-height: 16px; height: 62; width: 300;\" width=\"300\"></img></a>Giovanni Aldini may not be a household name, but his contributions to science cannot be ignored. Neither can his macabre demonstrations of the power of electricity on the human body.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/2012/05/15/mad-science-giovanni-aldini_n_1519723.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/2012/05/15/mad-science-giovanni-aldini_n_1519723.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(56)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "56",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[17]"
    },
    {
      "id": "-1181860611",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Obama Can't Knock the Hustle"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/robert-scheer/obama-cant-knock-the-hust_b_1523474.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Even after it was known that Jamie Dimon's bank blew more than $2 billion on the same suspect derivatives trading that has bankrupted the world's economy, Barack Obama still had praise for the intellect of his political backer and the integrity of the bank he heads."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1523474,\"mpid\":17}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/robert-scheer\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Robert Scheer\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/robert-scheer/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/robert-scheer/obama-cant-knock-the-hust_b_1523474.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">Obama Can't Knock the Hustle</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/robert-scheer\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Robert Scheer</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Editor, <a href=\"http://www.truthdig.com/\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\" target=\"_hplink\">Truthdig.com</a>; Author, 'The Great American Stickup'</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Even after it was known that Jamie Dimon's bank blew more than $2 billion on the same suspect derivatives trading that has bankrupted the world's economy, Barack Obama still had praise for the intellect of his political backer and the integrity of the bank he heads.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/robert-scheer/obama-cant-knock-the-hust_b_1523474.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/robert-scheer/obama-cant-knock-the-hust_b_1523474.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(78)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "78",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[18]"
    },
    {
      "id": "-808592255",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "What's Next for Facebook and Social Media? 5 Trends to Watch After the IPO"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/ryan-holmes/facebook-ipo_b_1522350.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "The Facebook IPO is a watershed moment in social media. It leaves no doubt that social networks are a true cultural and financial force."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1522350,\"mpid\":18}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/ryan-holmes\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Ryan Holmes\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/ryan-holmes/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/technology/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #3C3B41 !important; text-align: left; font: 12px Georgia;\">Tech</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/ryan-holmes/facebook-ipo_b_1522350.html\" style=\"outline: none; text-decoration: none; color: #3C3B41 !important;\">What's Next for Facebook and Social Media? 5 Trends to Watch After the IPO</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/ryan-holmes\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Ryan Holmes</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">CEO, HootSuite</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">The Facebook IPO is a watershed moment in social media.  It leaves no doubt that social networks are a true cultural and financial force. Social media is here to stay. It's not a fad.  And it's huge business. The big question is what's next.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/ryan-holmes/facebook-ipo_b_1522350.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/ryan-holmes/facebook-ipo_b_1522350.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[19]"
    },
    {
      "id": "168862312",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Fiat Ad Proves Once and for All That Breasts Can Sell Anything (VIDEO)"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/emma-gray/fiat-ad-proves-breasts-sell-anything-video_b_1519572.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "It's official -- breasts (and the prospect of breast augmentation) can sell anything. In case we weren't already convinced, a new Fiat ad has made it abundantly clear."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1519572,\"mpid\":19}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/emma-gray\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Emma Gray\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/emma-gray/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/women/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #BA72BA !important; text-align: left; font: 12px Georgia;\">Women</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/emma-gray/fiat-ad-proves-breasts-sell-anything-video_b_1519572.html\" style=\"outline: none; text-decoration: none; color: #BA72BA !important;\">Fiat Ad Proves Once and for All That Breasts Can Sell Anything (VIDEO)</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/emma-gray\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Emma Gray</a>, 05/16/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Associate Editor, HuffPost Women</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">It's official -- breasts (and the prospect of breast augmentation) can sell anything. In case we weren't already convinced, a new Fiat ad has made it abundantly clear.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/emma-gray/fiat-ad-proves-breasts-sell-anything-video_b_1519572.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/emma-gray/fiat-ad-proves-breasts-sell-anything-video_b_1519572.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(32)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "32",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[20]"
    },
    {
      "id": "-178268598",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "We Need Smarter Kids, Not More Smart Bombs"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/rev-jesse-jackson/we-need-smarter-kids-not_b_1524116.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Chicago is girding for the opening of the NATO Summit on May 20. The ministers and heads of state will be greeted by a rich array of protests, marches, events and counter-summits."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1524116,\"mpid\":20}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/rev-jesse-jackson\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Rev. Jesse Jackson\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/rev-jesse-jackson/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/rev-jesse-jackson/we-need-smarter-kids-not_b_1524116.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">We Need Smarter Kids, Not More Smart Bombs</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/rev-jesse-jackson\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Rev. Jesse Jackson</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Civil rights activist</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Chicago is girding for the opening of the NATO Summit on May 20. The ministers and heads of state will be greeted by a rich array of protests, marches, events and counter-summits. Security is already tight near the conference center, and tensions are building.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/rev-jesse-jackson/we-need-smarter-kids-not_b_1524116.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/rev-jesse-jackson/we-need-smarter-kids-not_b_1524116.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(44)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "44",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[21]"
    },
    {
      "id": "351589494",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "On Commencement Speeches"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/james-franco/on-commencement-speeches_b_1521338.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Commencement speeches are the worst kind of speech, because you need to be enthusiastic and inspiring in your own voice. There is nothing cheesier than that."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1521338,\"mpid\":21}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/james-franco\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"James Franco\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/james-franco/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/college/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #30577f !important; text-align: left; font: 12px Georgia;\">College</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/james-franco/on-commencement-speeches_b_1521338.html\" style=\"outline: none; text-decoration: none; color: #30577f !important;\">On Commencement Speeches</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/james-franco\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">James Franco</a>, 05/16/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Actor and author</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Commencement speeches are the worst kind of speech, because you need to be enthusiastic and inspiring in your own voice. There is nothing cheesier than that.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/james-franco/on-commencement-speeches_b_1521338.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/james-franco/on-commencement-speeches_b_1521338.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(70)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "70",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[22]"
    },
    {
      "id": "2050734786",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "5.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Why There's Nothing Wrong With a Bride Drinking Beer"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/natasha-burton/theres-nothing-wrong-with_b_1521651.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I don't like the idea that -- despite how the roles of women have changed -- a bride is assumed to fall in line with antiquated constraints on her wedding day. She might not be a virgin, but for the wedding she needs to appear chaste."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1521651,\"mpid\":22}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/natasha-burton\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Natasha Burton\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/natasha-burton/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/weddings/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #e2afc5; text-align: left; font: 12px Georgia;\">Weddings</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/natasha-burton/theres-nothing-wrong-with_b_1521651.html\" style=\"outline: none; text-decoration: none; color: #e2afc5;\">Why There's Nothing Wrong With a Bride Drinking Beer</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/natasha-burton\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Natasha Burton</a>, 05/16/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Staff Writer, HuffPost Weddings and HuffPost Divorce</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I don't like the idea that -- despite how the roles of women have changed -- a bride is assumed to fall in line with antiquated constraints on her wedding day. She might not be a virgin, but for the wedding she needs to appear chaste.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/natasha-burton/theres-nothing-wrong-with_b_1521651.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/natasha-burton/theres-nothing-wrong-with_b_1521651.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(126)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "126",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[23]"
    },
    {
      "id": "-1739909851",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Comparing Political Extremism in Europe and America"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/mugambi-jouet/comparing-political-extre_b_1506709.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "People often ask whether America or Europe fare better or worse in terms of political extremism. Yet, in a race to the bottom, what matters most is the direction."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1506709,\"mpid\":23}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/mugambi-jouet\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Mugambi Jouet\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/mugambi-jouet/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/mugambi-jouet/comparing-political-extre_b_1506709.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">Comparing Political Extremism in Europe and America</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/mugambi-jouet\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Mugambi Jouet</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Author</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">People often ask whether America or Europe fare better or worse in terms of political extremism. Yet, in a race to the bottom, what matters most is the direction.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/mugambi-jouet/comparing-political-extre_b_1506709.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/mugambi-jouet/comparing-political-extre_b_1506709.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\"></span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[24]"
    },
    {
      "id": "1264289762",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "5.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Could You Live Without Technology?"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/melissa-t-shultz/could-you-live-without-technology_b_1518758.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I've met other kids like my son who, unlike my generation, can disconnect quicker -- and not just because their parents or teachers have taken their cellphones and laptops away. I think they genuinely know how to hang out and have a good time without them."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1518758,\"mpid\":24}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/melissa-t-shultz\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Melissa T. Shultz\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/melissa-t-shultz/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/fifty/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #7e5bff; text-align: left; font: 12px Georgia;\">Post50</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/melissa-t-shultz/could-you-live-without-technology_b_1518758.html\" style=\"outline: none; text-decoration: none; color: #7e5bff;\">Could You Live Without Technology?</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/melissa-t-shultz\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Melissa T. Shultz</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Writer and Editor</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I've met other kids like my son who, unlike my generation, can disconnect quicker -- and not just because their parents or teachers have taken their cellphones and laptops away. I think they genuinely know how to hang out and have a good time without them. Could it be a trend?</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/melissa-t-shultz/could-you-live-without-technology_b_1518758.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/melissa-t-shultz/could-you-live-without-technology_b_1518758.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(17)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "17",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[25]"
    },
    {
      "id": "-821005187",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Wall Street, Romney, and Obama"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/mike-lux/romney-bain-capital_b_1522235.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Bain Capital was Wall Street at its worst. But the Obama team, in the White House and in the campaign, in order to win on the Bain attack, needs to face -- and turn around -- the perception that the administration has been weak on Wall Street."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1522235,\"mpid\":25}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/mike-lux\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Mike Lux\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/mike-lux/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/politics/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #03497E !important; text-align: left; font: 12px Georgia;\">Politics</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/mike-lux/romney-bain-capital_b_1522235.html\" style=\"outline: none; text-decoration: none; color: #03497E !important;\">Wall Street, Romney, and Obama</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/mike-lux\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Mike Lux</a>, 05/16/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Partner, Democracy Partners</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">Bain Capital was Wall Street at its worst. But the Obama team, in the White House and in the campaign, in order to win on the Bain attack, needs to face -- and turn around -- the perception that the administration has been weak on Wall Street.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/mike-lux/romney-bain-capital_b_1522235.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/mike-lux/romney-bain-capital_b_1522235.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(71)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "71",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[26]"
    },
    {
      "id": "1098030125",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "The New Wave of Underwater Hotels (Photos)"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/ben-hellwarth/the-new-underwater_hotels_b_1522460.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "I found a variety of new subsea suites that are on their way to opening their doors -- or hatches -- to vacationers who have deeper notions of what an ocean view should be."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"featBlg\",\"plid\":1522460,\"mpid\":26}}\"  style=\"color: #000; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; border: none !important; text-align: left; font: 12px Georgia;\">\n            <div style=\"width: 80px; float: left;\">\n                                                <a href=\"http://www.huffingtonpost.com/ben-hellwarth\" style=\"color: #058B7B; outline: none; text-decoration: none; text-align: left; font: 12px Georgia;\"><img alt=\"Ben Hellwarth\" height=\"45\" longdesc=\"http://s.huffpost.com/contributors/ben-hellwarth/headshot.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 45; width: 45;\" width=\"45\"></img></a>\n                \n                \n            </div>\n            <div style=\"width: 80px; float: left;\">\n            \t            \n            \t            \t<a data-beacon=\"{\"p\":{\"lnid\":\"vert\"}}\" href=\"http://www.huffingtonpost.com/travel/\" style=\"outline: none; text-decoration: none; font-weight: bold !important; font-family: Georgia; font-size: 10px; text-transform: uppercase; color: #12778E !important; text-align: left; font: 12px Georgia;\">Travel</a>\n            \t\n                <h3 style=\"color: #000; margin-bottom: 5px; line-height: 16px; text-align: left; font: 12px Georgia;\"><a href=\"http://www.huffingtonpost.com/ben-hellwarth/the-new-underwater_hotels_b_1522460.html\" style=\"outline: none; text-decoration: none; color: #12778E !important;\">The New Wave of Underwater Hotels (Photos)</a></h3>\n                <div style=\"color: #000; text-align: left; font: 12px Georgia;\"></div>\n\t\t\t</div>\n\t\t\t\t\t\t<div style=\"font-size: 1px; height: 1px !important; line-height: 1px !important; overflow: hidden !important; clear: both;\"></div>\n\t\t\t\n            <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\"><a href=\"http://www.huffingtonpost.com/ben-hellwarth\" style=\"outline: none; text-decoration: none; color: #111; font-style: italic; font-weight: bold; text-align: left; font: 12px Georgia;\">Ben Hellwarth</a>, 05/17/12</p>\n            <p style=\"margin-bottom: 8px; font-size: 11px !important; font-style: italic !important; line-height: 11px !important;\">Author and Journalist</p>\n                                <p style=\"margin-bottom: 8px; font-size: 13px; line-height: 16px;\">I found a variety of new subsea suites that are on their way to opening their doors -- or hatches -- to vacationers who have deeper notions of what an ocean view should be.</p>\n            \n            \n            <div>\n                <p style=\"color: #000; margin-bottom: 8px; font-size: 13px; line-height: 16px; text-align: left; font: 12px Georgia;\">\n                    <a href=\"http://www.huffingtonpost.com/ben-hellwarth/the-new-underwater_hotels_b_1522460.html\" style=\"outline: none; text-decoration: none; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%; color: #000 !important; font-weight: bold;\">Read Post</a>\n                    | <a href=\"http://www.huffingtonpost.com/ben-hellwarth/the-new-underwater_hotels_b_1522460.html#comments\" style=\"outline: none; text-decoration: none; color: #444; font-family: Arial, 'Helvetica Neue', Helvetica, sans-serif; font-size: 99%;\">Comments <span style=\"line-height: 16px; text-align: left; font: 12px Georgia;\">(11)</span></a>\n                    </p>\n                </div>\n            </div>"
          ]
        }
      ],
      "commentCount": "11",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[8]/DIV[1]/DIV[5]/DIV[4]/DIV[1]/DIV[1]/DIV[1]/DIV[2]/DIV[27]"
    },
    {
      "id": "429130710",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Russell Brand's Surprising Confession About Katy Perry"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/katy-perry-divorce-russell_n_1524652.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Obamas Divorce?.. Lies You Tell Kids.. Tiger Woods' Ex Single.. Two Magic"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524652,\"mpid\":1}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/katy-perry-divorce-russell_n_1524652.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Russell Brand's Surprising Confession About Katy Perry</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/katy-perry-divorce-russell_n_1524652.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Russell Brand\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610532/thumbs/s-RUSSELL-BRAND-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/katy-perry-divorce-russell_n_1524652.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/divorce\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #003b48; float: left;\">More Divorce</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #003b48; font-weight: bold; float: left;\"></div>Obamas Divorce?.. Lies You Tell Kids.. Tiger Woods' Ex Single.. Two Magic Words.. Hilarious E-Cards</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/katy-perry-divorce-russell_n_1524652.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/katy-perry-divorce-russell_n_1524652.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(139)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/video\",\"tag_name\":\"Video\", \"root_tag_id\":\"3\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Video</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "139",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[1]"
    },
    {
      "id": "1330859477",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "More Details Emerge In 'Terminator' Actor's Disappearance"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Hayden's Sexy Look.. TSwift's Donation.. Lea & Cory's PDA.. Snooki's"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524541,\"mpid\":2}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">More Details Emerge In 'Terminator' Actor's Disappearance</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Nick Stahl\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609608/thumbs/s-NICK-STAHL-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/celebrity\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #760101;\">More Celebrity</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #760101;\"></div>Hayden's Sexy Look.. TSwift's Donation.. Lea &amp; Cory's PDA.. Snooki's Sweet Tweet.. Jen's New Ad</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/nick-stahl-terminator-3-missing-in-la_n_1524541.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(86)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/@primacy_movies\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">Movies</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "86",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[2]"
    },
    {
      "id": "-1115802025",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Lindsay Lohan's Too Much Makeup.. Kirsten Dunst Child Star.. Smoothie Ingredient In Hair?"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.stylelist.com:80/2012/05/14/celebrity-sheer-clothes_n_1500578.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Nearly Sheer Katy Perry More Stylelist"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1500578,\"mpid\":3}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.stylelist.com/2012/05/14/celebrity-sheer-clothes_n_1500578.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\">Nearly Sheer</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.stylelist.com/2012/05/14/celebrity-sheer-clothes_n_1500578.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\" target=\"_hplink\"><img alt=\"Katy Perry\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610606/thumbs/s-KATY-PERRY-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.stylelist.com/2012/05/14/celebrity-sheer-clothes_n_1500578.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\"></a><a href=\"http://www.stylelist.com/\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #111;\">More Stylelist</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #111;\"></div>Lindsay Lohan's Too Much Makeup.. Kirsten Dunst Child Star..  Smoothie Ingredient In Hair?</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.stylelist.com/2012/05/14/celebrity-sheer-clothes_n_1500578.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.stylelist.com/2012/05/14/celebrity-sheer-clothes_n_1500578.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(26)</span></a>\n\t</li>\n  \n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523457,\"mpid\":4}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.ca/2012/05/17/jenna-talackova-miss-universe-canada-contestant-transgender_n_1523457.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Transgender Beauty Pageant Contestant Makes History</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.ca/2012/05/17/jenna-talackova-miss-universe-canada-contestant-transgender_n_1523457.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Jenna Talackova Miss Universe Canada\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609520/thumbs/s-JENNA-TALACKOVA-MISS-UNIVERSE-CANADA-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.ca/2012/05/17/jenna-talackova-miss-universe-canada-contestant-transgender_n_1523457.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.ca/2012/05/17/jenna-talackova-miss-universe-canada-contestant-transgender_n_1523457.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(515)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/jenna-talackova\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">Jenna Talackova</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "26",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[3],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[4]"
    },
    {
      "id": "2051074490",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "The Rick Santorum Joke That Didn't Make The WHCD"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/jimmy-kimmel-cut-rick-santorum-joke-white-house-correspondents-dinner_n_1524770.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Bill Maher's New Rule.. Marc Maron On Conan.. 'Anchorman 2' Poster.."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsb\",\"plid\":1524770,\"mpid\":5}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n    \n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/jimmy-kimmel-cut-rick-santorum-joke-white-house-correspondents-dinner_n_1524770.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">The Rick Santorum Joke That Didn't Make The WHCD</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/jimmy-kimmel-cut-rick-santorum-joke-white-house-correspondents-dinner_n_1524770.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Jimmy Kimmel\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610591/thumbs/s-JIMMY-KIMMEL-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/jimmy-kimmel-cut-rick-santorum-joke-white-house-correspondents-dinner_n_1524770.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/comedy\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #941300;\">More Comedy</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #941300;\"></div>Bill Maher's New Rule.. Marc Maron On Conan.. 'Anchorman 2' Poster.. Romney's New Slogans</a></h4>\n\n\n  \t\n  \n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/jimmy-kimmel-cut-rick-santorum-joke-white-house-correspondents-dinner_n_1524770.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/jimmy-kimmel-cut-rick-santorum-joke-white-house-correspondents-dinner_n_1524770.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\"></span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/jimmy-kimmel\",\"tag_name\":\"Jimmy Kimmel\", \"root_tag_id\":\"17682\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/jimmy-kimmel\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/jimmy-kimmel\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Jimmy Kimmel</a>\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[5]"
    },
    {
      "id": "-1555218154",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Sick Of Facebook? Try These 10 Alternatives Instead"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/16/facebook-alternatives_n_1522647.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Facebook Alternative Comments "
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1522647,\"mpid\":6}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/facebook-alternatives_n_1522647.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Sick Of Facebook? Try These 10 Alternatives Instead</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/facebook-alternatives_n_1522647.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Facebook Alternative\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609161/thumbs/s-FACEBOOK-ALTERNATIVE-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/facebook-alternatives_n_1522647.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/facebook-alternatives_n_1522647.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\"></span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/facebook\",\"tag_name\":\"Facebook\", \"root_tag_id\":\"1714\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/facebook\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/facebook\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Facebook</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "type": "LINK",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[6]"
    },
    {
      "id": "-1351159409",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "PHOTOS: The Most Daring Dress Ever ... No, Seriously"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "'Hair Vajazzling'.. When High-Slit Dresses Backfire.. Is This Kim Kardashian's"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1521069,\"mpid\":7}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">PHOTOS: The Most Daring Dress Ever ... No, Seriously</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Micaela Schaefer\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/608381/thumbs/s-MICAELA-SCHAEFER-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/style\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #7B0340;\">More Style</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #7B0340;\"></div>'Hair Vajazzling'.. When High-Slit Dresses Backfire.. Is This Kim Kardashian's Sexier Twin?</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/micaela-schaefer-tape-dress_n_1521069.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(2,127)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/fashion\",\"tag_name\":\"Fashion\", \"root_tag_id\":\"308\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/fashion\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/fashion\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Fashion</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "2",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[7]"
    },
    {
      "id": "1015383078",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "WATCH: Robotic Body Part Will Astound And Terrify You"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/robot-butt-shiri_n_1524229.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Amazing Little Film Narrated By Tom Waits.. 'Trick' Art.. Huge Marilyn"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524229,\"mpid\":8}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/robot-butt-shiri_n_1524229.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">WATCH: Robotic Body Part Will Astound And Terrify You</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/robot-butt-shiri_n_1524229.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Robotic Butt\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610215/thumbs/s-ROBOTIC-BUTT-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/robot-butt-shiri_n_1524229.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/culture\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #50B7AC;\">More Culture</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #50B7AC;\"></div>Amazing Little Film Narrated By Tom Waits.. 'Trick' Art.. Huge Marilyn Monroe Statue</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/robot-butt-shiri_n_1524229.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/robot-butt-shiri_n_1524229.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(143)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/smarter-ideas\",\"tag_name\":\"Smarter Ideas\", \"root_tag_id\":\"1171746\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/smarter-ideas\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/smarter-ideas\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Smarter Ideas</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "143",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[8]"
    },
    {
      "id": "-1130185666",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "GREATEST PERSON OF THE DAYMan Buys Out Entire Kmart Store, Donates It All To Charity"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/rankin-paynter-greatest-person_n_1523818.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Fight Famine.. Snooki's Surprise.. Cash For Organ Donors.. 'Man With"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523818,\"mpid\":9}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/rankin-paynter-greatest-person_n_1523818.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #ed0978; float: left;\">GREATEST PERSON OF THE DAY</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; float: left; height: 17px; width: 17px; background-color: #ed0978; font-weight: bold;\"></div>Man Buys Out Entire Kmart Store, Donates It All To Charity</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/rankin-paynter-greatest-person_n_1523818.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Greatest Person Of The Day\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610526/thumbs/s-GREATEST-PERSON-OF-THE-DAY-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/rankin-paynter-greatest-person_n_1523818.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/rankin-paynter-greatest-person_n_1523818.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/impact\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #ed0978; float: left;\">More Impact</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #ed0978; font-weight: bold; float: left;\"></div>Fight Famine.. Snooki's Surprise.. Cash For Organ Donors.. 'Man With Golden Voice' Returns</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/rankin-paynter-greatest-person_n_1523818.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/rankin-paynter-greatest-person_n_1523818.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(196)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/greatest-person-of-the-day\",\"tag_name\":\"Greatest Person Of The Day\", \"root_tag_id\":\"1285191\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/greatest-person-of-the-day\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/greatest-person-of-the-day\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Greatest Person Of The Day</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "196",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[9]"
    },
    {
      "id": "-953134265",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "WATCH: Thunder Star Gets Revenge Against World Peace"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/james-harden-elbows-metta-world-peace-video_n_1523719.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Preakness Odds.. Drew Brees Frustrated.. Rays Player Collapses.. Penalty"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523719,\"mpid\":10}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/james-harden-elbows-metta-world-peace-video_n_1523719.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">WATCH: Thunder Star Gets Revenge Against World Peace</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/james-harden-elbows-metta-world-peace-video_n_1523719.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"James Harden Elbow\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609586/thumbs/s-JAMES-HARDEN-ELBOW-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/james-harden-elbows-metta-world-peace-video_n_1523719.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/sports\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #363636;\">More Sports</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #363636;\"></div>Preakness Odds.. Drew Brees Frustrated.. Rays Player Collapses.. Penalty Box Fail</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/james-harden-elbows-metta-world-peace-video_n_1523719.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/james-harden-elbows-metta-world-peace-video_n_1523719.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(193)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/nba\",\"tag_name\":\"NBA\", \"root_tag_id\":\"19697\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/nba\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/nba\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">NBA</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "193",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[10]"
    },
    {
      "id": "130413937",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "The Naked Women Who Defied Hitler"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/judith-walkowitz/windmill-theatre-history_b_1522007.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "'50 Shades' Condo For Sale.. Famous Authors' Near-Death Experiences."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1522007,\"mpid\":11}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">The Naked Women Who Defied Hitler</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Windmill Theatre History\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609048/thumbs/s-WINDMILL-THEATRE-HISTORY-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/books\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #E61405; float: left;\">More Books</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #E61405; font-weight: bold; float: left;\"></div>'50 Shades' Condo For Sale.. Famous Authors' Near-Death Experiences.. Most Well-Read Cities </a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"><a data-beacon=\"{\"p\":{\"lnid\":\"read\"}}\" href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Read Post</a> |</li>\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/judith-walkowitz/windmill-theatre-history_b_1522007.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(114)</span></a>\n\t</li>\n  \n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "114",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[11]"
    },
    {
      "id": "-37605253",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "7 Things Women Really Want From Sex"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/16/sex-life-satisfaction-americans-fitness-magazine-yahoo-shine-survey_n_1522417.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "'Hysteria' Director Talks Vibrators.. Reasons To Love Paul Rudd.. Most"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1522417,\"mpid\":12}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/sex-life-satisfaction-americans-fitness-magazine-yahoo-shine-survey_n_1522417.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">7 Things Women Really Want From Sex</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/sex-life-satisfaction-americans-fitness-magazine-yahoo-shine-survey_n_1522417.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Sex Addict Pic\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/588023/thumbs/s-SEX-ADDICT-PIC-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/sex-life-satisfaction-americans-fitness-magazine-yahoo-shine-survey_n_1522417.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/women\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #BA72BA; float: left;\">More Women</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #BA72BA; font-weight: bold; float: left;\"></div>'Hysteria' Director Talks Vibrators.. Reasons To Love Paul Rudd.. Most Sexist Car Ad?</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/sex-life-satisfaction-americans-fitness-magazine-yahoo-shine-survey_n_1522417.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/sex-life-satisfaction-americans-fitness-magazine-yahoo-shine-survey_n_1522417.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(445)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/mostpopular\",\"tag_name\":\"Most Popular\", \"root_tag_id\":\"1382986\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/mostpopular\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/mostpopular\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Most Popular</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "445",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[12]"
    },
    {
      "id": "-1962470984",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Is It Really More Expensive To Eat Healthy?"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/16/healthy-food-expensive-junk-food-cost_n_1521882.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Healthy Food Expensive ... Comments (184)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1521882,\"mpid\":13}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n            <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/healthy-food-expensive-junk-food-cost_n_1521882.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Is It <em style=\"font-style: italic !important; font-size: 16px; line-height: 20px; font-weight: bold;\">Really</em> More Expensive To Eat Healthy?</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/healthy-food-expensive-junk-food-cost_n_1521882.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Healthy Food Expensive\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/608759/thumbs/s-HEALTHY-FOOD-EXPENSIVE-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n    \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/healthy-food-expensive-junk-food-cost_n_1521882.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/healthy-food-expensive-junk-food-cost_n_1521882.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(184)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/video\",\"tag_name\":\"Video\", \"root_tag_id\":\"3\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Video</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "184",
      "type": "LINK",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[13]"
    },
    {
      "id": "-1631590857",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "7 Ways To Avoid The #1 Reason For Divorce"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "No More Strapless Gowns.. End Your Affair NOW.. Cute Rainy-Day Weddings.."
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1521232,\"mpid\":14}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">7 Ways To Avoid The #1 Reason For Divorce</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Split Bills\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/608626/thumbs/s-SPLIT-BILLS-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: 5px 0;\"></a></div></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/weddings\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #e2afc5; float: left;\">More Weddings</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #e2afc5; font-weight: bold; float: left;\"></div>No More Strapless Gowns.. End Your Affair NOW.. Cute Rainy-Day Weddings.. Runaway Brides</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/avoiding-marriages-no-1-p_n_1521232.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\"></span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/after-the-wedding\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">After The Wedding</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[14]"
    },
    {
      "id": "-1745429613",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Mom Gets 9 Years In Teen's Malnutrition Death"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/angela-norman-sentenced-makayla-norman-malnutrition-neglect-cerebral-palsydeath_n_1524213.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "LISTEN: 911 Call.. Murder Investigation Reward.. Man Sewed Son's Buttocks"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524213,\"mpid\":15}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n            <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/angela-norman-sentenced-makayla-norman-malnutrition-neglect-cerebral-palsydeath_n_1524213.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Mom Gets 9 Years In Teen's Malnutrition Death</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/angela-norman-sentenced-makayla-norman-malnutrition-neglect-cerebral-palsydeath_n_1524213.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Angela Norman\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610396/thumbs/s-ANGELA-NORMAN-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/angela-norman-sentenced-makayla-norman-malnutrition-neglect-cerebral-palsydeath_n_1524213.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/crime\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #7f593e; float: left;\">More Crime</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #7f593e; font-weight: bold; float: left;\"></div>LISTEN: 911 Call.. Murder Investigation Reward.. Man Sewed Son's Buttocks Shut.. Dumb Criminals</a></h4>\n    \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/angela-norman-sentenced-makayla-norman-malnutrition-neglect-cerebral-palsydeath_n_1524213.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/angela-norman-sentenced-makayla-norman-malnutrition-neglect-cerebral-palsydeath_n_1524213.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(44)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/video\",\"tag_name\":\"Video\", \"root_tag_id\":\"3\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Video</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "44",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[2]/DIV[6]/DIV[3]/DIV[15]"
    },
    {
      "id": "-176534171",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "LISTEN: Singer's Greatest Hits.."
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Donna Summer Dies After Long Battle With Cancer ... Donna Summer Dead ... Family 'At Peace'.. "
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524410,\"mpid\":2}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Donna Summer Dies After Long Battle With Cancer</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Donna Summer Dead\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610459/thumbs/s-DONNA-SUMMER-DEAD-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-songs-dead-video_n_1524442.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">LISTEN: Singer's Greatest Hits.. </a> <a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Family 'At Peace'.. </a> <a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-dies-twitter_n_1524674.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Celebrities React.. </a> <a href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dies-style-evolution_n_1524677.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">A Look At Her Dazzling Style</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/donna-summer-dead-queen-of-disco-dies_n_1524410.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(1,488)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/video\",\"tag_name\":\"Video\", \"root_tag_id\":\"3\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Video</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "1",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[2]"
    },
    {
      "id": "-571867965",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "White House Aims To Curb Prison Rape"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/prison-rape-elimination-act-immigrant-detention_n_1524470.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Bill Seeks To Expand Anti-Sexual Assault Provisions To Immigrant Detention Centers"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1525186,\"mpid\":3}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n            <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/prisons-step-up-anti-rape_n_1525186.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">White House Aims To Curb Prison Rape</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/prisons-step-up-anti-rape_n_1525186.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Prison Rape Justice\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610824/thumbs/s-PRISON-RAPE-JUSTICE-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/prisons-step-up-anti-rape_n_1525186.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/2012/05/17/prison-rape-elimination-act-immigrant-detention_n_1524470.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\">Bill Seeks To Expand Anti-Sexual Assault Provisions To Immigrant Detention Centers</a></h4>\n    \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/prisons-step-up-anti-rape_n_1525186.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/prisons-step-up-anti-rape_n_1525186.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(95)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/obama-administration\",\"tag_name\":\"Obama Administration\", \"root_tag_id\":\"170524\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/obama-administration\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/obama-administration\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Obama Administration</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "95",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[3]"
    },
    {
      "id": "1122606878",
      "sp": "0.100",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Supreme Court Health Care Ruling Could Mean Life Or Death For Some"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/supreme-court-health-care_n_1525120.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Birther Rant.. Mitt, Rick Agree On Pee.. GOPer vs. 'Mobs'.. Tea Party"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1525120,\"mpid\":4}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n            <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/supreme-court-health-care_n_1525120.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Supreme Court Health Care Ruling Could Mean Life Or Death For Some</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/supreme-court-health-care_n_1525120.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Supreme Court Health Care\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610745/thumbs/s-SUPREME-COURT-HEALTH-CARE-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/supreme-court-health-care_n_1525120.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/politics\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #03497E;\">More Politics</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #03497E;\"></div>Birther Rant.. Mitt, Rick Agree On Pee.. GOPer vs. 'Mobs'.. Tea Party Nightmare.. Anti-Walker Weapon</a></h4>\n    \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/supreme-court-health-care_n_1525120.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/supreme-court-health-care_n_1525120.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(67)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/health-care\",\"tag_name\":\"Health Care\", \"root_tag_id\":\"597\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/health-care\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/health-care\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Health Care</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523253,\"mpid\":5}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/pinterest-funding_n_1523253.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">$1.5 BILLION?</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/pinterest-funding_n_1523253.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Pinterest\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/532930/thumbs/s-PINTEREST-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/pinterest-funding_n_1523253.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/technology\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #3C3B41;\">More Tech</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #3C3B41;\"></div>Apple Censorship.. Facebook Drinking Game.. Google Car... Comcast Fee.. Google Update</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/pinterest-funding_n_1523253.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/pinterest-funding_n_1523253.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(135)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/pinterest\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">Pinterest</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524778,\"mpid\":6}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/bob-marshall-virginia-gop-gay-judge_n_1524778.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Lawmaker Defends Opposition To Gay Judge: 'Sodomy Is Not A Civil Right'</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/bob-marshall-virginia-gop-gay-judge_n_1524778.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Bob Marshall Tracy Thorne Begland\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610634/thumbs/s-BOB-MARSHALL-TRACY-THORNE-BEGLAND-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/bob-marshall-virginia-gop-gay-judge_n_1524778.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/bob-marshall-virginia-gop-gay-judge_n_1524778.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(949)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/video\",\"tag_name\":\"Video\", \"root_tag_id\":\"3\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Video</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "67",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[4],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[5],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[6]"
    },
    {
      "id": "-124321590",
      "sp": "-0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Reporters Prevented From Asking Biden Questions At Campaign Event"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/joe-biden-reporters-blocked-ohio_n_1524672.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Buffett's Huge Buy.. Rush Makes Trouble.. Amazing Maddow.. Bill's Trayvon"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524672,\"mpid\":7}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/joe-biden-reporters-blocked-ohio_n_1524672.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Reporters Prevented From Asking Biden Questions At Campaign Event</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/joe-biden-reporters-blocked-ohio_n_1524672.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Joe Biden\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610609/thumbs/s-JOE-BIDEN-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/joe-biden-reporters-blocked-ohio_n_1524672.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/media\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #DD5200;\">More Media</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #DD5200;\"></div>Buffett's Huge Buy.. Rush Makes Trouble.. Amazing Maddow.. Bill's Trayvon Attack.. Mitt Freaks!</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/joe-biden-reporters-blocked-ohio_n_1524672.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/joe-biden-reporters-blocked-ohio_n_1524672.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(245)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/joe-biden\",\"tag_name\":\"Joe Biden \", \"root_tag_id\":\"628\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/joe-biden\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/joe-biden\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Joe Biden </a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1525213,\"mpid\":8}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n                        <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://nymag.com/daily/intel/2012/05/gop-health-care-plan-still-due-any-day-now.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\">GOP's Health Care Plan Remains A Mirage</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://nymag.com/daily/intel/2012/05/gop-health-care-plan-still-due-any-day-now.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\" target=\"_hplink\"><img alt=\"Gop Health Care Plan\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610795/thumbs/s-GOP-HEALTH-CARE-PLAN-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n             \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/gop-health-care-plan-still-due-any-day_n_1525213.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/gop-health-care-plan-still-due-any-day_n_1525213.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(50)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/health-care\",\"tag_name\":\"Health Care\", \"root_tag_id\":\"597\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/health-care\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/health-care\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Health Care</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524940,\"mpid\":9}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n            <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/ilya-yashin-russia-protests_n_1524940.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Russian Opposition Leader Jailed</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/ilya-yashin-russia-protests_n_1524940.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Ilya Yashin\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610675/thumbs/s-ILYA-YASHIN-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n    \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/ilya-yashin-russia-protests_n_1524940.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/ilya-yashin-russia-protests_n_1524940.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(12)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/vladimir-putin\",\"tag_name\":\"Vladimir Putin\", \"root_tag_id\":\"3779\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/vladimir-putin\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/vladimir-putin\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Vladimir Putin</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "245",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[7],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[8],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[9]"
    },
    {
      "id": "-821993543",
      "sp": "-0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "The People Who Still Haven't Joined Facebook"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/facebook-users-resisters-holdouts_n_1523349.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Comments (487)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523349,\"mpid\":10}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/facebook-users-resisters-holdouts_n_1523349.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">The People Who Still Haven't Joined Facebook</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/facebook-users-resisters-holdouts_n_1523349.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Facebook\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/607238/thumbs/s-FACEBOOK-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/facebook-users-resisters-holdouts_n_1523349.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: 5px 0;\"></a></div></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/facebook-users-resisters-holdouts_n_1523349.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/facebook-users-resisters-holdouts_n_1523349.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(487)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/ap\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">AP</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1525221,\"mpid\":11}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/mitt-romney-ricketts-plan_n_1525221.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Romney: 'I Stand By What I Said, Whatever It Was'</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/mitt-romney-ricketts-plan_n_1525221.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Mitt Romney Ricketts Plan\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610804/thumbs/s-MITT-ROMNEY-RICKETTS-PLAN-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/mitt-romney-ricketts-plan_n_1525221.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/mitt-romney-ricketts-plan_n_1525221.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(31)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/mitt-romney-2012\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">Mitt Romney 2012</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1525031,\"mpid\":12}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n<div style=\"margin-bottom: 5px;\"><a href=\"http://www.huffingtonpost.com/jason-linkins\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #03497E;\">HuffPost Reports</div><div style=\"background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; float: left; background-color: #03497E;\"></div><span style=\"position: relative; top: 1px; font-weight: bold !important; font-family: Arial; font-size: 12px; text-transform: uppercase; color: #03497E !important;\">Jason Linkins</span></a></div>\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/americans-elect_n_1525031.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Americans Elect Continues To Fail At The Internet</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/americans-elect_n_1525031.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Americanselect\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610701/thumbs/s-AMERICANSELECT-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/americans-elect_n_1525031.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/americans-elect_n_1525031.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\"></span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/2012\",\"tag_name\":\"2012\", \"root_tag_id\":\"90886\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/2012\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/2012\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">2012</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "487",
      "type": "LINK",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[10],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[11],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[12]"
    },
    {
      "id": "-813314449",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Jon Stewart vs. Fox News Pundits On 'The Daily Show'"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/jon-stewart-vs-fox-news-pundits-on-the-daily-show-video_n_1519343.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Jon Stewart Fox News ... Comments (461)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1519343,\"mpid\":13}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/jon-stewart-vs-fox-news-pundits-on-the-daily-show-video_n_1519343.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Jon Stewart vs. Fox News Pundits On 'The Daily Show'</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/jon-stewart-vs-fox-news-pundits-on-the-daily-show-video_n_1519343.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Jon Stewart Fox News\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/607430/thumbs/s-JON-STEWART-FOX-NEWS-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/jon-stewart-vs-fox-news-pundits-on-the-daily-show-video_n_1519343.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/jon-stewart-vs-fox-news-pundits-on-the-daily-show-video_n_1519343.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/jon-stewart-vs-fox-news-pundits-on-the-daily-show-video_n_1519343.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(461)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/fox-news\",\"tag_name\":\"Fox News\", \"root_tag_id\":\"6526\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/fox-news\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/fox-news\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Fox News</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "461",
      "type": "LINK",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[13]"
    },
    {
      "id": "532016689",
      "sp": "0.100",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "GOP Being 'Really Immature. Irresponsible. Let's Get Serious'"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/nancy-pelosi-john-boehner-debt-limit_n_1524466.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Pelosi, Boehner Trade Barbs Over Debt Limit Fight ... Nancy Pelosi John Boehner Debt Limit ... Comments (1,235)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524466,\"mpid\":14}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nancy-pelosi-john-boehner-debt-limit_n_1524466.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Pelosi, Boehner Trade Barbs Over Debt Limit Fight</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nancy-pelosi-john-boehner-debt-limit_n_1524466.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Nancy Pelosi John Boehner Debt Limit\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610541/thumbs/s-NANCY-PELOSI-JOHN-BOEHNER-DEBT-LIMIT-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/nancy-pelosi-john-boehner-debt-limit_n_1524466.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">GOP Being 'Really Immature. Irresponsible. Let's Get Serious'</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/nancy-pelosi-john-boehner-debt-limit_n_1524466.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/nancy-pelosi-john-boehner-debt-limit_n_1524466.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(1,235)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/deficit\",\"tag_name\":\"Deficit\", \"root_tag_id\":\"48291\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/deficit\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/deficit\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Deficit</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524257,\"mpid\":15}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/rand-paul-war-iran-syria-sanctions_n_1524257.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Amendment Barring War With Iran Added To Senate Bill </a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/rand-paul-war-iran-syria-sanctions_n_1524257.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Rand Paul\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/599376/thumbs/s-RAND-PAUL-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/rand-paul-war-iran-syria-sanctions_n_1524257.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/rand-paul-war-iran-syria-sanctions_n_1524257.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(543)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/syria\",\"tag_name\":\"Syria\", \"root_tag_id\":\"1002\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/syria\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/syria\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Syria</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524268,\"mpid\":16}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/fbi-nypd-robert-mueller-ray-kelly-al-qaeda-terror-plot-chuck-schumer_n_1524268.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">FBI Didn't Tell NYPD About Terror Plot</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/fbi-nypd-robert-mueller-ray-kelly-al-qaeda-terror-plot-chuck-schumer_n_1524268.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Robert Mueller Fbi Nypd Terror\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610223/thumbs/s-ROBERT-MUELLER-FBI-NYPD-TERROR-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/fbi-nypd-robert-mueller-ray-kelly-al-qaeda-terror-plot-chuck-schumer_n_1524268.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/fbi-nypd-robert-mueller-ray-kelly-al-qaeda-terror-plot-chuck-schumer_n_1524268.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(40)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/al-qaeda\",\"tag_name\":\"Al Qaeda\", \"root_tag_id\":\"2184\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/al-qaeda\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/al-qaeda\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Al Qaeda</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "1",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[14],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[15],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[16]"
    },
    {
      "id": "-1170157795",
      "sp": "0.100",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "WATCH: Footage Reportedly Shows UN Observers Under Fire In Syria"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/un-observer-idlib_n_1524529.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Obama Campaign Calls 'B.S.' On Karl Rove Roveobama Comments (1,735)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523874,\"mpid\":17}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/obama-campaign-karl-rove-crossroads-gps-ad_n_1523874.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Obama Campaign Calls 'B.S.' On Karl Rove</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/obama-campaign-karl-rove-crossroads-gps-ad_n_1523874.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Roveobama\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610245/thumbs/s-ROVEOBAMA-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/obama-campaign-karl-rove-crossroads-gps-ad_n_1523874.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: 5px 0;\"></a><a href=\"http://www.huffingtonpost.com/2012/05/17/obama-campaign-karl-rove-crossroads-gps-ad_n_1523874.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/obama-campaign-karl-rove-crossroads-gps-ad_n_1523874.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/obama-campaign-karl-rove-crossroads-gps-ad_n_1523874.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(1,735)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/elections-2012\",\"tag_name\":\"Elections 2012\", \"root_tag_id\":\"1295155\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/elections-2012\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/elections-2012\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Elections 2012</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524529,\"mpid\":18}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/un-observer-idlib_n_1524529.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">WATCH: Footage Reportedly Shows UN Observers Under Fire In Syria</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/un-observer-idlib_n_1524529.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Idlib\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610523/thumbs/s-IDLIB-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/un-observer-idlib_n_1524529.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/un-observer-idlib_n_1524529.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(15)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/syria\",\"tag_name\":\"Syria\", \"root_tag_id\":\"1002\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/syria\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/syria\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Syria</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523906,\"mpid\":19}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/mary-kennedy-suicide-dead-dies_n_1523906.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Mary Kennedy Reportedly Hanged Herself</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/mary-kennedy-suicide-dead-dies_n_1523906.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Mary Kennedy Dies\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609636/thumbs/s-MARY-KENNEDY-DIES-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/mary-kennedy-suicide-dead-dies_n_1523906.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/mary-kennedy-suicide-dead-dies_n_1523906.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(148)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/mostpopular\",\"tag_name\":\"Most Popular\", \"root_tag_id\":\"1382986\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/mostpopular\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/mostpopular\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Most Popular</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "1",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[17],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[18],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[19]"
    },
    {
      "id": "1373338520",
      "sp": "0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "James O'Keefe Fails Again"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://www.huffingtonpost.com:80/2012/05/17/james-okeefe-voter-fraud-video_n_1524146.html"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "James Okeefe Voter Fraud ... Comments (437)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1524146,\"mpid\":20}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/james-okeefe-voter-fraud-video_n_1524146.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">James O'Keefe Fails <em style=\"font-style: italic !important; font-size: 16px; line-height: 20px; font-weight: bold;\">Again</em></a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/james-okeefe-voter-fraud-video_n_1524146.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"James Okeefe Voter Fraud\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/610274/thumbs/s-JAMES-OKEEFE-VOTER-FRAUD-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/james-okeefe-voter-fraud-video_n_1524146.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/james-okeefe-voter-fraud-video_n_1524146.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/james-okeefe-voter-fraud-video_n_1524146.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(437)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/video\",\"tag_name\":\"Video\", \"root_tag_id\":\"3\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/video\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Video</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "437",
      "type": "LINK",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[20]"
    },
    {
      "id": "576309490",
      "sp": "0.100",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Nobody Knows How Big Losses Will End Up Being"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://dealbook.nytimes.com:80/2012/05/16/jpmorgans-trading-loss-is-said-to-rise-at-least-50/"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "More JPMorgan Losses ... Jamie Dimon ... Comments (3,111)"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523074,\"mpid\":21}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n                        <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://dealbook.nytimes.com/2012/05/16/jpmorgans-trading-loss-is-said-to-rise-at-least-50/\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\"><em style=\"font-style: italic !important; font-size: 16px; line-height: 20px; font-weight: bold;\">More</em> JPMorgan Losses</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://dealbook.nytimes.com/2012/05/16/jpmorgans-trading-loss-is-said-to-rise-at-least-50/\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\" target=\"_hplink\"><img alt=\"Jamie Dimon\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609327/thumbs/s-JAMIE-DIMON-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://dealbook.nytimes.com/2012/05/16/jpmorgans-trading-loss-is-said-to-rise-at-least-50/\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\"></a><a href=\"http://www.huffingtonpost.com/2012/05/16/jpmorgan-chase-chief-investment-office_n_1520377.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Bank Played By Different High-Risk Rules.. </a><a href=\"http://www.reuters.com/article/2012/05/17/us-jpmorgan-losses-idUSBRE84G06T20120517\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Nobody Knows How Big Losses Will End Up Being</a></h4>\n             \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/jpmorgan-chase-trading-loss_n_1523074.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/jpmorgan-chase-trading-loss_n_1523074.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(3,111)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t<span hovercard_params=\"{\"tag_url\":\"/news/banks\",\"tag_name\":\"Banks\", \"root_tag_id\":\"18516\"}\"   style=\"font: 12px/16px Arial,'Helvetica Neue',Helvetica,sans-serif; font-weight: bold; margin: 0px !important; list-style-type: disc;\">\n\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/banks\"  style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; margin-right: 1px; padding: 1px; background: transparent url(http://s.huffpost.com:80/images/entries/follow-bignews.gif) no-repeat 50% 50%; height: 15px; width: 14px; display: inline-block; margin: 0px !important; text-align: left;\">&nbsp;</a>\n\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/banks\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; margin: 0px !important; text-align: left;\">Banks</a>\n\t\t\t\t\n\t\t\t</span>\n\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523147,\"mpid\":22}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/rep-mike-coffman-colorado_n_1523147.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">GOP Lawmaker: Obama Is 'Just Not An American'</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/rep-mike-coffman-colorado_n_1523147.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Mike Coffman Obama\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609320/thumbs/s-MIKE-COFFMAN-OBAMA-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/rep-mike-coffman-colorado_n_1523147.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/rep-mike-coffman-colorado_n_1523147.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/rep-mike-coffman-colorado_n_1523147.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(3,909)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/colorado\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">Colorado</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "3",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[21],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[22]"
    },
    {
      "id": "1441343575",
      "sp": "-0.000",
      "fresh": "1.000",
      "sr": "4.000",
      "tagName": "item",
      "cluster": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]",
      "childNodes": [
        {
          "tagName": "title",
          "childNodes": [
            "Trayvon Martin's Autopsy Results Released"
          ]
        },
        {
          "tagName": "link",
          "childNodes": [
            "http://usnews.msnbc.msn.com:80/_news/2012/05/16/11736208-trayvon-martin-killed-by-single-gunshot-fired-from-intermediate-range-autopsy-shows?lite"
          ]
        },
        {
          "tagName": "pubDate",
          "childNodes": [
            "Thu, 17 May 2012 20:06:26 GMT"
          ]
        },
        {
          "tagName": "textSummary",
          "childNodes": [
            "Home HIV Test.. Fight To Free O.J... Celebs Stand By Obama.. Fired Writer's"
          ]
        },
        {
          "tagName": "description",
          "childNodes": [
            "<div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523272,\"mpid\":23}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n                        <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://usnews.msnbc.msn.com/_news/2012/05/16/11736208-trayvon-martin-killed-by-single-gunshot-fired-from-intermediate-range-autopsy-shows?lite\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\">Trayvon Martin's Autopsy Results Released</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://usnews.msnbc.msn.com/_news/2012/05/16/11736208-trayvon-martin-killed-by-single-gunshot-fired-from-intermediate-range-autopsy-shows?lite\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\" target=\"_hplink\"><img alt=\"Trayvon Martin Plaque\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609340/thumbs/s-TRAYVON-MARTIN-PLAQUE-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://usnews.msnbc.msn.com/_news/2012/05/16/11736208-trayvon-martin-killed-by-single-gunshot-fired-from-intermediate-range-autopsy-shows?lite\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\" target=\"_hplink\"></a><a href=\"http://www.huffingtonpost.com/black-voices\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); float: left; background-color: #ED4A4B;\">More Black Voices</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; font-weight: bold; float: left; background-color: #ED4A4B;\"></div>Home HIV Test.. Fight To Free O.J... Celebs Stand By Obama.. Fired Writer's New Racist Rant</a></h4>\n             \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/trayvon-martin-autopsy_n_1523272.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/trayvon-martin-autopsy_n_1523272.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(2,155)</span></a>\n\t</li>\n  \n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\"> \n\t\t| \n\t\t\t\t\t\t\t\t\t\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"bigNws\"}}\" href=\"http://www.huffingtonpost.com/news/trayvon-martin\" style=\"outline: none; text-decoration: none; margin-right: 1px; padding: 1px; color: #0088c3; font-size: 12px; font-weight: bold; list-style-type: disc;\">Trayvon Martin</a>\t\t\t\n\t\t\t\t\n\t\t\n\t</li>\n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1521275,\"mpid\":24}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n    <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/gay-librarian-rejects-christian-lifestyle-statement_n_1521275.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Mass Exodus At Baptist University Over Anti-Gay 'Lifestyle Statement' </a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/gay-librarian-rejects-christian-lifestyle-statement_n_1521275.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Shorter University\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/608645/thumbs/s-SHORTER-UNIVERSITY-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a><div style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; bottom: 45px; float: left; height: 1px; left: 0; margin: 0; padding: 0; position: relative; text-align: left; z-index: 2;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/gay-librarian-rejects-christian-lifestyle-statement_n_1521275.html\" style=\"color: #058B7B; outline: none; text-decoration: none; background: url(http://s.huffpost.com:80/images/hoton_sprite.png) no-repeat left top; float: left; height: 45px; width: 45px; background-position: -50px 0px;\"></a></div></div><h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/16/gay-librarian-rejects-christian-lifestyle-statement_n_1521275.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"></a><a href=\"http://www.huffingtonpost.com/religion\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\"><div style=\"-moz-border-radius: 4px 0 0 4px; -webkit-border-radius: 4px 0 0 4px; border-radius: 4px 0 0 4px; line-height: 15px; padding: 1px 1px 1px 6px; font-weight: bold !important; font-family: Arial; font-size: 12px; color: #fff !important; text-transform: uppercase; -moz-box-shadow: 1px 3px 4px rgba(0,0,0,.09); -webkit-box-shadow: 1px 3px 4px rgba(0,0,0,.09); box-shadow: 1px 3px 4px rgba(0,0,0,.09); background-color: #771C85; float: left;\">More Religion</div><div style=\"font-size: 16px; line-height: 20px; background-image: url(http://s.huffpost.com:80/images/v/tag_arrow.png); background-position: 0px -17px; background-repeat: no-repeat; height: 17px; width: 17px; background-color: #771C85; font-weight: bold; float: left;\"></div>Syrian Christians.. Respect Everyone.. Music And Holiness.. Psalms And Evil</a></h4>\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/16/gay-librarian-rejects-christian-lifestyle-statement_n_1521275.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/16/gay-librarian-rejects-christian-lifestyle-statement_n_1521275.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(2,585)</span></a>\n\t</li>\n  \n\n</ul>\n\n\n\n\n</div><div data-beacon=\"{\"p\":{\"mnid\":\"newsa\",\"plid\":1523246,\"mpid\":25}}\"  style=\"color: #000; font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; border-bottom: 1px solid #c2c2c2; margin-top: 8px; padding-top: 8px; margin: 0; padding: 13px 0 12px; text-align: left;\">\n\n\n            <h4 style=\"font-size: 16px; line-height: 20px; margin-bottom: 4px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/ecuador-drug-lab-discover_n_1523246.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; outline: none; text-decoration: none; color: #1a1a1a; text-align: left;\">Police Make Shocking Discovery At Site Of Plane Crash</a></h4><div style=\"margin-bottom: 9px;\"><a href=\"http://www.huffingtonpost.com/2012/05/17/ecuador-drug-lab-discover_n_1523246.html\" style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #058B7B; outline: none; text-decoration: none; text-align: left;\"><img alt=\"Cocaine\" height=\"219\" longdesc=\"http://i.huffpost.com/gen/609333/thumbs/s-COCAINE-large300.jpg\" src=\"http://www.huffingtonpost.com/images/trans.gif\" style=\"height: 219; width: 300;\" width=\"300\"></img></a></div>\n    \n    \n\n\n\n  <ul>\n\t\n\n\t\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; display: none !important; text-align: left;\">http://www.huffingtonpost.com/2012/05/17/ecuador-drug-lab-discover_n_1523246.html</li>\n\n\n\t<li style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; color: #aaa; margin-right: 1px; text-align: left;\">\n\t\t<a data-beacon=\"{\"p\":{\"lnid\":\"cmnts\"}}\" href=\"http://www.huffingtonpost.com/2012/05/17/ecuador-drug-lab-discover_n_1523246.html#comments\" style=\"outline: none; text-decoration: none; color: #3f3f3f; font-size: 13px; font-weight: bold; margin-right: 1px; padding: 1px; list-style-type: disc;\">Comments <span style=\"font: 12px/16px Arial, 'Helvetica Neue', Helvetica, sans-serif; text-align: left;\">(301)</span></a>\n\t</li>\n  \n\n</ul>\n\n\n\n\n</div>"
          ]
        }
      ],
      "commentCount": "2",
      "type": "STORY",
      "xroot": "/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[23],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[24],/HTML[1]/BODY[1]/DIV[4]/DIV[3]/DIV[3]/DIV[9]/DIV[1]/DIV[2]/DIV[25]"
    }
  ]
}
`
