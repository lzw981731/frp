// Copyright 2017 fatedier, fatedier@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vhost

import (
	"bytes"
	"io/ioutil"
	"net/http"

	frpLog "github.com/fatedier/frp/utils/log"
	"github.com/fatedier/frp/utils/version"
)

var (
	NotFoundPagePath = ""
)

const (
	NotFound = `<!DOCTYPE html>
<html>
<head>
<title>Not Found</title>
<style>
    body {
        width: 35em;
        margin: 0 auto;
        font-family: Tahoma, Verdana, Arial, sans-serif;
    }
</style>
<head>
    <meta charset="utf-8" />
        <title>網站已遭查禁</title>
        <!-- Fonts -->
        <link  rel="stylesheet">
        <!-- CSS -->
        <link rel="stylesheet"  integrity="sha384-/Y6pD6FV/Vv2HJnA6t+vslU6fwYXjCFtcEpHbNJ0lyAFsXTsjBbfaDjzALeQsN6M" crossorigin="anonymous">
    <style type="text/css">
        body{
            font-family: 'Raleway', sans-serif;
            font-size: 14px;
            -webkit-font-smoothing: antialiased;
            }

            h1, h2, h3, h4, h5{
            font-family:'Montserrat', sans-serif;
            }
            
            .img-content{
            position: absolute;
            align-content: center;
            text-align: center;
            margin: auto;
            height: 50%;
            top: -50%;
            right: 0;
            bottom: 0;
            left: 0;
            }
            .content{
            position: absolute;
            align-content: center;
            text-align: center;
            margin: auto;
            height: 50%;
            top: 0;
            right: 0;
            bottom: 0;
            left: 0;
            }

            .banner{
            padding:3%;
            align-content: center ;
            margin: auto 5% auto 5% ;
            background-image: linear-gradient(to top, #30cfd0 0%, #330867 100%);
            border-radius: 20px;
            box-shadow: 3px 3px 5px 6px #cccccc;
            }

            .banner h1{
            font-family: "Microsoft JhengHei" ;
            font-size: 50px;   
            margin-top:1%;
            color: #FFF;
            }
        
            .banner h2{
            color: #FFF;
            margin-bottom:3%;
            font-size:40px;
            }

            .banner h3{
            color: #FFF;
            opacity: .9;
            font-size:25px;
            }

            .banner p{
            color: #FFF;
            font-size: 24px;
            padding: 20px 0;
            font-weight: 300;
            width: 70%;
            margin: 0 auto;
            }
    </style>
    <!--
<script async src='/cdn-cgi/bm/cv/2172558837/api.js'></script>
-->
</head>

<body>
<section class="bg-1">
    <div class="content">
        <div style="margin-top:-11%;">
            <img src="https://cdn.jsdelivr.net/gh/lzw981731/img/2020/06/26/5f5d.png" style="align-content: center;text-align: center;width: 40%;height: auto; margin: auto auto 1% auto; ">
        </div>
        <div class="banner">
                <div class="col-12 text-center">
                    <div style="background-color:firebrick;">
                        <h1>網站已經遭到查禁</h1>
                    </div>
                        <h2>(This Domain Has Been Seized)</h2>
                        <h3>已經違背中華民國著作權法第九十一條及九十二條規範，全部或部分內容涉屬盜版，正進入司法偵查中。</h3>
                        <p>
                        The website is in violation of the Copyright Act of the Republic of China (Taiwan) for its unauthorized use of materials.  Investigation is currently underway.</p>
                        <p>內政部警政署刑事警察局電偵大隊 敬啟 <br/>
                        Criminal Investigation Bureau（Taiwan）</p>
                </div>
            </div>
    </div>
</section>
<script type="text/javascript">(function(){window['__CF$cv$params']={r:'580bbb10ab56988d',m:'89a8a3de2b7eced49bf0d7d5400be20e69b51188-1586346206-1800-AXfMipt69xS9k39OseM0Cvbr0f3Kbs2a6i15NQFaHcJ2EE29f4nRPtpVQp/KgduXqsjmOfivLCxx7gGaor0X8kEbM7J50HQX16lmdjp4bzFe+Nr/xYV+BXL01+ME0Wcofw==',s:[0xde0ce21efa,0xa321d7537b],fb:0,}})();</script></body>
</html>
`
)

func getNotFoundPageContent() []byte {
	var (
		buf []byte
		err error
	)
	if NotFoundPagePath != "" {
		buf, err = ioutil.ReadFile(NotFoundPagePath)
		if err != nil {
			frpLog.Warn("read custom 404 page error: %v", err)
			buf = []byte(NotFound)
		}
	} else {
		buf = []byte(NotFound)
	}
	return buf
}

func notFoundResponse() *http.Response {
	header := make(http.Header)
	header.Set("server", "frp/"+version.Full())
	header.Set("Content-Type", "text/html")

	res := &http.Response{
		Status:     "Not Found",
		StatusCode: 404,
		Proto:      "HTTP/1.0",
		ProtoMajor: 1,
		ProtoMinor: 0,
		Header:     header,
		Body:       ioutil.NopCloser(bytes.NewReader(getNotFoundPageContent())),
	}
	return res
}

func noAuthResponse() *http.Response {
	header := make(map[string][]string)
	header["WWW-Authenticate"] = []string{`Basic realm="Restricted"`}
	res := &http.Response{
		Status:     "401 Not authorized",
		StatusCode: 401,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     header,
	}
	return res
}
