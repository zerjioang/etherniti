// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

// +build prod pre

package api

var (
	// bad bot blacklisted user agent strings
	badBotsList = []string{
		"almaden",
		"anarchie",
		"aspseek",
		"attach",
		"autoemailspider",
		"backweb",
		"bandit",
		"batchftp",
		"blackwidow",
		"bot",
		"buddy",
		"bumblebee",
		"cherrypicker",
		"chinaclaw",
		"cicc",
		"collector",
		"copier",
		"crescent",
		"custo",
		"da",
		"diibot",
		"disco",
		"disco pump",
		"download demon",
		"download wonder",
		"downloader",
		"drip",
		"dsurf15a",
		"ecatch",
		"easydl/2.99",
		"eirgrabber",
		"email",
		"emailcollector",
		"emailsiphon",
		"emailwolf",
		"express webpictures",
		"extractorpro",
		"eyenetie",
		"filehound",
		"flashget",
		"frontpage",
		"getright",
		"getsmart",
		"getweb!",
		"gigabaz",
		"go\\!zilla",
		"go!zilla",
		"go-ahead-got-it",
		"gotit",
		"grabber",
		"grabnet",
		"grafula",
		"grub-client",
		"googlebot",
		"hmview",
		"httrack",
		"httpdown",
		"httrack",
		"ia_archiver",
		"image stripper",
		"image sucker",
		"indy*library",
		"indy library",
		"interget",
		"internetlinkagent",
		"internet ninja",
		"internetseer.com",
		"iria",
		"jbh*agent",
		"jetcar",
		"joc web spider",
		"justview",
		"larbin",
		"leechftp",
		"lexibot",
		"lftp",
		"link*sleuth",
		"likse",
		"link",
		"linkwalker",
		"mag-net",
		"magnet",
		"masscan",
		"mass downloader",
		"memo",
		"microsoft.url",
		"midown tool",
		"mirror",
		"mister pix",
		"mozilla.*indy",
		"mozilla.*newt",
		"mozilla*msiecrawler",
		"ms frontpage*",
		"msfrontpage",
		"msiecrawler",
		"msproxy",
		"navroad",
		"nearsite",
		"netants",
		"netmechanic",
		"netspider",
		"net vampire",
		"netzip",
		"nicerspro",
		"ninja",
		"octopus",
		"offline explorer",
		"offline navigator",
		"openfind",
		"pagegrabber",
		"papa foto",
		"pavuk",
		"pcbrowser",
		"ping",
		"pingalink",
		"pockey",
		"psbot",
		"pump",
		"qrva",
		"realdownload",
		"reaper",
		"recorder",
		"reget",
		"scooter",
		"seeker",
		"siphon",
		"sitecheck.internetseer.com",
		"sitesnagger",
		"slysearch",
		"smartdownload",
		"snake",
		"spacebison",
		"sproose",
		"stripper",
		"sucker",
		"superbot",
		"superhttp",
		"surfbot",
		"szukacz",
		"takeout",
		"teleport pro",
		"urlspiderpro",
		"vacuum",
		"voideye",
		"web image collector",
		"web sucker",
		"webauto",
		"[ww]eb[bb]andit",
		"webcollage",
		"webcopier",
		"web downloader",
		"webemailextrac.*",
		"webfetch",
		"webgo is",
		"webhook",
		"webleacher",
		"webminer",
		"webmirror",
		"webreaper",
		"websauger",
		"website",
		"website extractor",
		"website quester",
		"webster",
		"webstripper",
		"webwhacker",
		"webzip",
		"whacker",
		"widow",
		"wwwoffle",
		"wow",
		"x-tractor",
		"xaldon webspider",
		"xenu",
		"zeus.*webster",
		"zeus",
		//literal strings
		"windows 95",
		"windows 98",
		"biz360.com",
		"xpymep",
		"turnitinbot",
		"sindice",
		"purebot",
		"wget",
		"libwww-perl",
		"apachebench",
		// "curl",        //allow in development and production mode
	}
)
