package gowiki

import (
	"encoding/json"
	"fmt"
	"strings"
)

const StyleCss = `
<style>
*{
  box-sizing: border-box
}

body {
  font-family: Arial, Helvetica, sans-serif;
  margin: 0;
} 
.header {
  padding: 20px;
  text-align: center;
  background: #39CCCC;
  color: #333333;
}
.header h1 {
  font-size: 40px;
  color: #333333;
}
.navbar {
  overflow: hidden;
  background-color: Navy
}
.navbar a {
  float: left;
  display: block;
  color:  white;   
  text-align: center;
  padding: 14px 20px;
  text-decoration: none
}
.navbar a.right {
  float: right
}
.navbar a:hover {
  background-color: #ddd;
  color: white;
}
.row {  
  display: flex;
  flex-wrap: wrap;
  height: 1200px;
}
.side {
  flex: 28%;
  background-color:  #f7f7f7 ;
  border: 1.5px solid #e9e9e9;
  padding: 10px;
}
.side pre{
  border-radius: 7px;
	padding: 1px;
  border: 1.9px solid  #e9e9e9 ;
  white-space: pre-wrap;
  overflow-wrap: break-word; 
  background-color: #f7f7f7;
  font-size: 12px;
  overflow: hidden;
}
.string { color: #cb4b16; }
.number { color: #859900; }
.boolean { color: blue; }
.null { color: magenta; }
.key { color: #456971; }

 .main {
  color: #333333;
  flex: 72%;
  background-color: white;
  padding-left: 20px;
  border: 1px solid  #e9e9e9 ;
}

.footer {
  padding: 10px;
  text-align: center;
  background:  #39CCCC;
}
  
.w3-container {
  padding:0.01em 16px;
  padding-right: 10px;
}

.w3-container:after, .w3-container:before {
  content:"";
  display:table;
  clear:both
}

.w3-table-all {
  border-collapse:collapse;
  border-spacing:0;
  width:90%;
  display:table;
  border:1px solid #ccc;
  padding-right: 10px;
}

.w3-table-all tr {
  border-bottom:1px solid #ddd;
  background-color:#fff;
  background-color:#f1f1f1
}

.w3-table-all td,.w3-table-all th {
  border:1px solid #ccc;
  padding:8px 8px;
  display:table-cell;
  text-align:left;
  vertical-align:top;
  padding-left:16px
}

.w3-card-4{
  box-shadow:0 4px 10px 0 rgba(0,0,0,0.2),0 4px 20px 0 rgba(0,0,0,0.19)
}

@media screen and (max-width: 700px) {
  .row {   
    flex-direction: column;
  }
}

@media screen and (max-width: 400px) {
  .navbar a {
    float: none;
    width:100%
  }
}
.btn {
  border-radius: 5px;
  border:1px solid black;
  color: white;
  padding: 14px 28px;
  font-size: 16px;
  cursor: pointer
}

.success {background-color: #4CAF50;}  
.success:hover {background-color: #46a049;}

.info {background-color: #2196F3;}  
.info:hover {background: #0b7dda;}

.warning {background-color: #ff9800;}  
.warning:hover {background: #e68a00;}

.danger {background-color: #f44336;}   
.danger:hover {background: #da190b;}

.default {background-color: #e7e7e7; color: black; font-size: 10px;padding: auto; width=20px; height=20px;border-radius: 25px;}   
.default:hover {background: #ddd;}
</style>`

var Main = `<!DOCTYPE html>
<html lang="en">

<head>
    <title>Pagina de Alonso</title>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    STYLECSS
</head>

<body>

    <div class="header">
        <h1>Controlador Local</h1>
        <p>Una interfaz de software privado</p>
    </div>

    <div class="navbar">
        <a href="/masinfo/">Saber mas</a>
        <a href="/sobreme/" class="right">Local</a>
    </div>

    <div class="row">
        <div class="side">
        <h3>Ultimo evento emitido por el sensor</h3>
          <p Id="example"></p>
          <form action="/main/" method="GET">
          <button class="btn default">F5</button>
          </form>
        </div>
        <div class="main">
            <h2>Estado en la Base de Datos</h2>
            <h4>Id Sensor: {{.IdEquip}}; Sector: {{.Sector}}; Usuario: {{.User}}; Estado: {{.Status}};</h4>
            <h2>Control Ventilador</h2>
            <h4>API: Particle Sensores, Abr 1, 2020</h4>
            <button class="btn success">Encender Ventilador</button>
            <button class="btn success">Apagar Ventilador</button>
            <br>
            <h2>Control Quemador y Llama</h2>
            <button class="btn warning">Encender Quemador</button>
            <button class="btn success">Apagar Quemador</button>
            <button class="btn warning">Encender Llama</button>
            <button class="btn success">Apagar Llama</button>
            <button class="btn danger">Apagar Todo</button>
            <h5></h5>
            <h2>Tabla de eventos</h2>
            <h4>Tabla de Ultimos estados publicados</h4>
            <button class="btn info">Actualizar Tabla</button>
            <br>
            <br>
            <table class="w3-table-all w3-card-4">
                <tbody>
                    <tr>
                        <th>A</th>
                        <th>B</th>
                        <th>C</th>
                        <th>D</th>
                        <th>E</th>
                    </tr>
                    <tr>
                        <td>12.2</td>
                        <td>12.12</td>
                        <td>Encendido</td>
                        <td>2020-03-19T20:12:21</td>
                        <td>Chile</td>
                    </tr>
                    <tr>
                        <td>12.2</td>
                        <td>12.12</td>
                        <td>Encendido</td>
                        <td>2020-03-19T20:12:21</td>
                        <td>Chile</td>
                    </tr>
                    <tr>
                        <td>12.2</td>
                        <td>12.12</td>
                        <td>Encendido</td>
                        <td>2020-03-19T20:12:21</td>
                        <td>Chile</td>
                    </tr>
                    <tr>
                        <td>12.2</td>
                        <td>12.12</td>
                        <td>Encendido</td>
                        <td>2020-03-19T20:12:21</td>
                        <td>Chile</td>
                    </tr>
                    <tr>
                        <td>12.2</td>
                        <td>12.12</td>
                        <td>Encendido</td>
                        <td>2020-03-19T20:12:21</td>
                        <td>Chile</td>
                    </tr>

                </tbody>
            </table>
        </div>
    </div>

    <div class="footer">
        <h2>*-*</h2>
    </div>
    <script type="text/javascript">
    function output(inp) {
      x = document.getElementById("example")
      x.innerHTML= "<pre>"+inp+"</pre>"
      //document.body.appendChild(document.createElement('pre')).innerHTML = inp;
    }

    function syntaxHighlight(json) {
        json = json.replace(/&/g, '&amp;').replace(/</g, '&lt;').replace(/>/g, '&gt;');
        return json.replace(/("(\\u[a-zA-Z0-9]{4}|\\[^u]|[^\\"])*"(\s*:)?|\b(true|false|null)\b|-?\d+(?:\.\d*)?(?:[eE][+\-]?\d+)?)/g, function (match) {
            var cls = 'number';
            if (/^"/.test(match)) {
                if (/:$/.test(match)) {
                    cls = 'key';
                } else {
                    cls = 'string';
                }
            } else if (/true|false/.test(match)) {
                cls = 'boolean';
            } else if (/null/.test(match)) {
                cls = 'null';
            }
            return '<span class="' + cls + '">' + match + '</span>';
        });
    }
    </script>
    <script>
      var obj = {{.Side}};
      //console.log({{.Side}});
      //console.log(obj)
      obj=JSON.parse(obj)
      var str = JSON.stringify(obj, undefined, 4);
      output(syntaxHighlight(str));
    </script>
    <script>
      // tell the embed parent frame the height of the content
      if (window.parent && window.parent.parent) {
          window.parent.parent.postMessage(["resultsFrame", {
              height: document.body.getBoundingClientRect().height,
              slug: ""
          }], "*")
      }
      // always overwrite window.name, in case users try to set it manually
      window.name = "result"
    </script>


</body>

</html>
`

const Pi = "<!DOCTYPE html><h1>Hola mundo</h1><h1>Hola mundo</h1><h1>Hola mundo</h1><h1>Hola mundo</h1>"

const Edit = `<h1>Editing {{.Title}}</h1>
<form action="/save/{{.Title}}" method="POST">
<div><textarea name="body" rows="20" cols="80">{{printf "%s" .Body}}</textarea></div>
<div><input type="submit" value="Save"></div>
</form>
`

//QUIERO GRABAR CAMBIAR SOLO ESE VALOR

func FusionJsonHtmlEstatus(a interface{}) string {
	formatjson, _ := json.MarshalIndent(a, "", "	")
	//fmt.Printf("%+v\n", string(formatjson))
	load := ` <pre><code>Y</code></pre>`
	new := strings.Replace(load, "Y", string(formatjson), 10000)
	fmt.Println(new)
	return strings.Replace(Main, "", new, 10000)
}
