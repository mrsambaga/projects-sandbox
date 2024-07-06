import { calculateInvestmentResults, formatter } from "../util/investment";

export default function ResultTable({ investmentInput }) {
  const annualData = calculateInvestmentResults(investmentInput);
  let initialInvestment = 0;
  if (annualData.length > 0) {
    initialInvestment =
      annualData[0].valueEndOfYear -
      annualData[0].interest -
      annualData[0].annualInvestment;
  }

  return (
    <table id="result">
      <thead>
        <tr>
          <th>Year</th>
          <th>Investment Value</th>
          <th>Interest (Year)</th>
          <th>Total Interest</th>
          <th>Invested Capital</th>
        </tr>
      </thead>
      <tbody>
        {annualData.map((data, index) => (
          <tr key={index}>
            <td>{data.year}</td>
            <td>{formatter.format(data.valueEndOfYear)}</td>
            <td>{formatter.format(data.interest)}</td>
            <td>
              {formatter.format(
                data.valueEndOfYear -
                  data.annualInvestment * data.year -
                  initialInvestment
              )}
            </td>
            <td>
              {formatter.format(
                data.valueEndOfYear -
                  data.valueEndOfYear +
                  data.annualInvestment * data.year +
                  initialInvestment
              )}
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
}
