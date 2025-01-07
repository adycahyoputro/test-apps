// $("#save").click(function () {
//   var obj = {};
//   obj.username = $("#exampleInputText1").val();
//   obj.email = $("#exampleInputEmail1").val();
//   var myJson = JSON.stringify(obj);
//   console.log(myJson);
//   $.ajax({
//     url: "http://192.168.157.160:8082/register",
//     type: "POST",
//     contentType: "application/json",
//     data: myJson,
//     success: function (result) {
//       // console.log(result);
//       console.log(result.data.error);
//       // console.log(result.error);
//       // alert("success");
//       if (result.data.error != "") {
//         alert("success");
//       } else {
//         alert("failed");
//       }
//     },
//   });
// });
var x = document.getElementById("save");
x.addEventListener("click", AddData);

async function AddData() {
  var obj = {};
  obj.username = $("#exampleInputText1").val();
  obj.email = $("#exampleInputEmail1").val();
  var data = JSON.stringify(obj);
  // try {
  let endpoint = "http://192.168.157.160:8082/register";
  const request = await fetch(endpoint, {
    method: "POST",
    // body: JSON.stringify(obj),
    body: data,
    headers: { "Content-Type": "application/json" },
  });

  const response = await request.json();
  // console.log(response);

  if (response.data) {
    alert("success register");
    console.log(response);
  }
  if (response.error) {
    alert("failed");
    console.log(response);
  }
}
