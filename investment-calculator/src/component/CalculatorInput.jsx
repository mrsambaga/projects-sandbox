import { useState } from "react";

export default function CalculatorInput({
  inputLabel,
  onChangeInput,
  inputName,
}) {
  const [inputValue, setInputValue] = useState("");

  function handleInputChange(event) {
    setInputValue(event.target.value);
    onChangeInput(event);
  }

  return (
    <div id="user-input">
      <label>{inputLabel}</label>
      <input
        value={inputValue}
        onChange={handleInputChange}
        name={inputName}
        type="number"
      />
    </div>
  );
}
