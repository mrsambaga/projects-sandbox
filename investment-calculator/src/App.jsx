import { useState } from "react";
import logo from "./assets/investment-calculator-logo.png";
import CalculatorInput from "./component/CalculatorInput";
import ResultTable from "./component/ResultTable";

function App() {
  const [investmentInput, setInvestmentInput] = useState({
    initialInvestment: 0,
    annualInvestment: 0,
    expectedReturn: 0,
    duration: 0,
  });

  function onInputChange(event) {
    const { name, valueAsNumber } = event.target;
    const newInput = {
      ...investmentInput,
      [name]: valueAsNumber,
    };
    setInvestmentInput(newInput);
  }

  return (
    <>
      <div id="header">
        <img src={logo}></img>
        <h1>React Investment Calculator</h1>
      </div>
      <div className="input-group">
        <CalculatorInput
          inputLabel="INITIAL INVESTMENT"
          inputName="initialInvestment"
          onChangeInput={onInputChange}
        ></CalculatorInput>
        <CalculatorInput
          inputLabel="ANNUAL INVESTMENT"
          inputName="annualInvestment"
          onChangeInput={onInputChange}
        ></CalculatorInput>
        <CalculatorInput
          inputLabel="EXPECTED RETURN"
          inputName="expectedReturn"
          onChangeInput={onInputChange}
        ></CalculatorInput>
        <CalculatorInput
          inputLabel="DURATION"
          inputName="duration"
          onChangeInput={onInputChange}
        ></CalculatorInput>
      </div>
      <div>
        <ResultTable investmentInput={investmentInput} />
      </div>
    </>
  );
}

export default App;
