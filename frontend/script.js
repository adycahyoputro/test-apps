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
//       console.log(result);
//   if (result.data != "") {
//     alert("success");
//   } else {
//     alert("failed");
//   }
//     },
//   });
// });
var x = document.getElementById("save");
x.addEventListener("click", AddData);

var obj = {};
obj.username = $("#exampleInputText1").val();
obj.email = $("#exampleInputEmail1").val();
async function AddData() {
  let endpoint = "http://192.168.157.160:8082/register";
  const request = await fetch(endpoint, {
    method: "POST",
    //   body: data,
    body: JSON.stringify(obj),
    headers: { "Content-Type": "application/json" },
  });

  const response = await request.json();
  console.log(response);
  //   if (response.ok) {
  //     console.log(response);
  //   } else {
  //     console.log(`An error with status code ${response.status} occured`);
  //   }
}
AddData();
