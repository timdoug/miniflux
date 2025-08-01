// SPDX-FileCopyrightText: Copyright The Miniflux Authors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package rewrite // import "miniflux.app/v2/internal/reader/rewrite"

// List of predefined rewrite rules (alphabetically sorted)
// domain => rule name
var predefinedRules = map[string]string{
	"abstrusegoose.com":      "add_image_title",
	"amazingsuperpowers.com": "add_image_title",
	"blog.cloudflare.com":    `add_image_title,remove("figure.kg-image-card figure.kg-image + img")`,
	"cowbirdsinlove.com":     "add_image_title",
	"drawingboardcomic.com":  "add_image_title",
	"exocomics.com":          "add_image_title",
	"framatube.org":          "nl2br,convert_text_link",
	"happletea.com":          "add_image_title",
	"ilpost.it":              `remove(".art_tag, #audioPlayerArticle, .author-container, .caption, .ilpostShare, .lastRecents, #mc_embed_signup, .outbrain_inread, p:has(.leggi-anche), .youtube-overlay")`,
	"imogenquest.net":        "add_image_title",
	"lukesurl.com":           "add_image_title",
	"medium.com":             "fix_medium_images",
	"mercworks.net":          "add_image_title",
	"monkeyuser.com":         "add_image_title",
	"mrlovenstein.com":       "add_image_title",
	"nedroid.com":            "add_image_title",
	"oglaf.com":              `replace("media.oglaf.com/story/tt(.+).gif"|"media.oglaf.com/comic/$1.jpg"),add_image_title`,
	"optipess.com":           "add_image_title",
	"peebleslab.com":         "add_image_title",
	"quantamagazine.org":     `add_youtube_video_from_id, remove("h6:not(.byline,.post__title__kicker), #comments, .next-post__content, .footer__section, figure .outer--content, script")`,
	"qwantz.com":             "add_image_title,add_mailto_subject",
	"sentfromthemoon.com":    "add_image_title",
	"thedoghousediaries.com": "add_image_title",
	"theverge.com":           `add_dynamic_image, remove("div.duet--recirculation--related-list, .hidden")`,
	"treelobsters.com":       "add_image_title",
	"xkcd.com":               "add_image_title",
}
