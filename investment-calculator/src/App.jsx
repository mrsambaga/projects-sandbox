import { useState } from "react";
import logo from "./assets/investment-calculator-logo.png";
import CalculatorInput from "./component/CalculatorInput";
import ResultTable from "./component/ResultTable";
import { calculateInvestmentResults } from "./util/investment";

function App() {
  const [annualData, setAnnualData] = useState([]);
  const [investmentInput, setInvestmentInput] = useState({
    initialInvestment: 0,
    annualInvestment: 0,
    expectedReturn: 0,
    duration: 0,
  });

  const newAnnualData = calculateInvestmentResults(investmentInput);

  function onInputChange(event) {
    const { name, value } = event.target;
    const newInput = {
      ...investmentInput,
      [name]: value,
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
        <ResultTable></ResultTable>
      </div>
    </>
  );
}

export default App;
