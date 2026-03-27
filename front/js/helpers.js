function convertFields(data, numberFields = []) {
  const result = { ...data };

  for (const field of numberFields) {
    if (field in result && result[field] !== "") {
      result[field] = Number(result[field]);
    }
  }

  return result;
}