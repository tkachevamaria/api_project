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
// Получение

async function getEmployee() {
  const id = document.getElementById("getId").value;

  const res = await fetch(`${API_URL}/employee/${id}`);

  const result = await res.json();
  document.getElementById("result").textContent =
    JSON.stringify(result, null, 2);
}

//-------------------
//Обновление
// const form = document.getElementById("update-form");
// form.addEventListener("submit", (e) => {
//   e.preventDefault();  // отменяем стандартное поведение формы
//   updateEmployee(form);     // вызываем функцию для отправки обновления
// });

async function updateEmployee(id, data){
  const res = await fetch(`${API_URL}/employee/${id}`, {
    method: "PUT",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });
  const result = await res.json();
  document.getElementById("result").textContent =
    JSON.stringify(result, null, 2);
}

//------------------------------
// универсальная функция подключения форм
function attachForms() {
  
  const forms = document.querySelectorAll("form[data-action]");

  forms.forEach(form => {
    form.addEventListener("submit", async (e) => {
      e.preventDefault();
      console.log("submit работает");

      const rawData = Object.fromEntries(new FormData(form).entries());
      const action = form.dataset.action;

      switch (action){
      case "create": {
        const data = convertFields(rawData, ["age", "salary"]);
        await createEmployee(data);
      }

      case "update": {
        const id = rawData.id;
        if (!id) return alert("ты id забыл придурок");
        delete rawData.id;
        const data = convertFields(rawData, ["age", "salary"]);
        await updateEmployee(id, data);
      }
    }
    });
  });
}

// подключаем все формы сразу
attachForms();