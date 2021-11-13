let mypain = document.getElementById("2");
let val = document.getElementById("3").value;
mypain.onclick = function(){
  var url = "http://localhost:8080/add";
  var invocation = new XMLHttpRequest();
  var body = val;
  function callOtherDomain(){
    if(invocation)
      {
        invocation.open("POST", url, true);
        invocation.setRequestHeader('X-PINGOTHER', 'pingpong');
        invocation.setRequestHeader('Content-Type', 'application/xml');
        invocation.onreadystatechange = handler;
        invocation.send(val);
      }
  }
}