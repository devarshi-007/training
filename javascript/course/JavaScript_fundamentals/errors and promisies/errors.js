export default function () {
  try {
    let car = j;
  } catch (error) {
    console.log("error: ", error.message, error.stack);

    try {
      throw new Error("custom error");
    } catch {
      console.log("handled");
    }
  } finally {
    console.log("end...");
  }

  console.log("continuing...");
}
