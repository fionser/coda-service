<!DOCTYPE html>
<html>
<head>
<title>Beego</title>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.12.4/css/bootstrap-select.min.css">
<script src="https://code.jquery.com/jquery-2.2.4.min.js" integrity="sha256-BbhdlvQf/xTY9gja0Dq3HiwQF8LaCRTXxZKRutelT44=" crossorigin="anonymous"></script>
<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
</head>
<body>
<ul class="nav nav-tabs">
<li role="presentation"><a onclick='show_ctnr("new_session_ctnr");' href="#new">New Session</a></li>
<li role="presentation"><a onclick='show_ctnr("gen_key_ctnr");' href="#gen">Generate Key-pair</a></li>
<li role="presentation"><a onclick='show_ctnr("enc_ctnr");' href="#enc">Encrypt Data</a></li>
</ul>

<form action="/v1/session" method="POST" enctype="multipart/form-data">
<div class = "row" id = "new_session_ctnr">
<div class = "col-lg-6">
<div class="input-group from-group">
    <input type="text" class="form-control" name="userName" placeholder="User Name" aria-describedby="sizing-addon2">
    <span class="input-group-addon" id="basic-addon2">Protocol</span>
    <select id='protocol-list' name="protocol">
    </select>
</div>
    <input type="text" class="form-control" name="sessionName" placeholder="Session Name" aria-describedby="sizing-addon2">
    <input type="file" class="form-control" name="schemaFile" aria-describedby="sizing-addon2">
    <input type="text" class="form-control" name="clientList" placeholder="Clients" aria-describedby="sizing-addon2">
    <input type='submit'>
</div>
</div>
</form>

<div class="row hidden"  id="gen_key_ctnr">
<div class="col-lg-6">
<div class="input-group">
<input type="text" class="form-control" placeholder="User Name" aria-describedby="sizing-addon2">
<input type="text" class="form-control" placeholder="Session Name" aria-describedby="sizing-addon2">
</div>
</div>
</div>

<div class="row hidden"  id="enc_ctnr">
<div class="col-lg-6">
<div class="input-group">
<input type="text" class="form-control" placeholder="User Name" aria-describedby="sizing-addon2">
<input type="text" class="form-control" placeholder="Session Name" aria-describedby="sizing-addon2">
<input type="file" class="form-control" title="Data">
</div>
</div>
</div>


<script>
    var activated_ctnr = 'new_session_ctnr';
    var containers = ["new_session_ctnr", "gen_key_ctnr", "enc_ctnr"];
    var rootURL = '/v1';
    function show_ctnr(ID) {
        if (containers.indexOf(ID) <  0) {
            return;
        }

        if (ID != activated_ctnr) {
            $('#' + activated_ctnr).addClass("hidden");
            $('#' + ID).removeClass("hidden");
            activated_ctnr = ID;
        }
    }

    function getProtocols() {
        var list = $('#protocol-list');
        $.ajax({
               type: 'GET',
               url: rootURL + "/protocol",
               dataType: "json",
               success: function(protocols) {
                   list.empty();
                   for (var idx in protocols) {
                       var protocol = protocols[idx];
                       var option = $('<option/>').text(protocol.Name).appendTo(list);
                   }
                   $('select').selectpicker();
               }
               });
    }
    getProtocols();
</script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/bootstrap-select/1.12.4/js/bootstrap-select.min.js"></script>
</body>
</html>
