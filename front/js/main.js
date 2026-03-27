const API_URL = "http://localhost:8080";

// --------------------
// СОЗДАНИЕ
async function createEmployee() {
  const data = {
    name: document.getElementById("name").value,
    sex: document.getElementById("sex").value,
    age: Number(document.getElementById("age").value),
    salary: Number(document.getElementById("salary").value)
  };

  //отправка post запроса
  const res = await fetch(`${API_URL}/employee`, {
    //await - чтобы подождать пока выполнится fetch
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    //преобразование в json
    body: JSON.stringify(data)
  });

  const result = await res.json();
  document.getElementById("result").textContent =
    JSON.stringify(result, null, 2);
}

// --------------------
// ПОЛУЧЕНИЕ
async function getEmployee() {
  const id = document.getElementById("getId").value;

  const res = await fetch(`${API_URL}/employee/${id}`);

  const result = await res.json();
  document.getElementById("result").textContent =
    JSON.stringify(result, null, 2);
}