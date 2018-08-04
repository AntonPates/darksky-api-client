package converter

var summaryTpl = `
<html>
<head>
<title>Weather forecast</title>
<meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
<meta charset="UTF-8">
</head>
<body>
<style type="text/css">
      canvas {
        display: block;
        float: left;
        margin-left: 8px;
        margin-right: 8px;
        margin-top: 20px;
      }
</style>
{{range $k,$v := .}}
<a name="{{.Location.Anchor}}"></a>
	<h2>Погода {{.Location.InName}}</h2>
  <h3>Сейчас {{.Location.InName}}</h3>
  <!-- <div style="width: 100%; height: 75px; background-image: url('{{.Currently.Icon}}'); background-repeat: no-repeat; padding-left: 75px;"> -->
  <div><canvas id="{{.Location.Anchor}}-{{$k}}" width="64" height="64"></canvas>
  Температура {{.Currently.Temperature | TempColor}} C (ощущается как {{.Currently.ApparentTemperature | TempColor}} C), {{.Currently.Summary | ToLower}}<br>
  Влажность {{.Currently.Humidity | Humidity}}%<br>
  Давление {{.Currently.Pressure | Pressure}} мм. ртутного столба<br>
  Ветер {{.Currently.WindBearing | WindDirection}}, {{.Currently.WindSpeed  | WindSpeed}} м/с.
  </div>
<div style="clear:both;"> </div>
<details>  
<summary style="cursor: pointer; font-size: large;">Прогноз погоды {{.Location.InName}} на 2 суток по часам</summary>
<div style="width: 656px; margin-left: -5px; display: block;">
  {{range $key,$val := .Hourly.Data}}
{{if future .Time }}
{{if and (gt $key 0) (mod $key 2)}} <div style="clear:both;"> </div> {{end}}  
<div style="width: 325px; float: left;">
<div><strong>Прогноз на {{.Time | TimeFunc}}</strong></div>
   <canvas id="{{$v.Location.Anchor}}-{{$k}}-hourly-{{$key}}" width="64" height="64"></canvas>
    <div>
    {{.Temperature | TempColor}} C, {{.Summary | ToLower}} <br>
    Влажность {{.Humidity | Humidity}}%<br>
    Давление {{.Pressure | Pressure}} мм. ртутного столба<br>
    Ветер {{.WindBearing | WindDirection}}, {{.WindSpeed | WindSpeed}} м/с.</div>
    </div>
{{end}}
{{end}}
</div>
</details>
<div style="clear:both;"> </div>

<h3>Прогноз погоды {{.Location.InName}} на 7 дней</h3>
<div style="width: 656px; margin-left: -5px; display: block;">
  {{range $key,$val := .Daily.Data}}
{{if future .Time }}
{{if and (gt $key 0) (mod $key 2)}} <div style="clear:both;"> </div> {{end}}  
<div style="width: 325px; float: left;">
<div><strong>Прогноз на {{.Time | DateFunc}}</strong></div>
   <canvas id="{{$v.Location.Anchor}}-{{$k}}-daily-{{$key}}" width="64" height="64"></canvas>
    <div>
    Днем {{.TemperatureHigh | TempColor}} C, {{.Summary | ToLower}} <br>
    Ночью {{.TemperatureLow | TempColor}} C.<br>
    Ветер {{.WindBearing | WindDirection}}, {{.WindSpeed | WindSpeed}} м/с.</div>
    </div>
{{end}}    
  {{end}}
</div>
<div style="clear:both;"> </div>
{{end}}

 <script src="skycons.js"></script>
<script>
      var skyicons = new Skycons();
{{range $k,$v := .}}
  skyicons.add("{{.Location.Anchor}}-{{$k}}", Skycons.{{.Currently.Icon | IconDescription}});
  {{range $key,$val := .Hourly.Data}}
  skyicons.add("{{$v.Location.Anchor}}-{{$k}}-hourly-{{$key}}", Skycons.{{.Icon | IconDescription}});
   {{end}}
  {{range $key,$val := .Daily.Data}}
     skyicons.add("{{$v.Location.Anchor}}-{{$k}}-daily-{{$key}}", Skycons.{{.Icon | IconDescription}});
  {{end}}
{{end}}
      skyicons.play();
    </script>
</body>
</html>
`

var currentlyTpl = `
<h1 class="widget-title">Погода на Байкале сейчас</h1>
{{range .}}
<a href="/weather-on-baikal/#{{.Location.Anchor}}">{{.Location.Name}}</a> {{.Currently.Temperature | TempColor}} C, {{.Currently.Summary | ToLower}}<br />
{{end}}
<div style="height: 15px;">&nbsp;</div>
`
