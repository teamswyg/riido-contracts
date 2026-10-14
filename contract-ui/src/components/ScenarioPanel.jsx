export default function ScenarioPanel({ operation }) {
  const scenarios = operation.scenarios ?? [];
  if (scenarios.length === 0) return null;
  return (
    <section className="scenario-panel">
      <h3>Figma and contract scenarios</h3>
      <div className="scenario-list">
        {scenarios.map((scenario) => (
          <article className="scenario" key={scenario.name}>
            <h4>{scenario.name}</h4>
            <p><span>Given</span>{scenario.given}</p>
            <p><span>When</span>{scenario.when}</p>
            <p><span>Then</span>{scenario.then}</p>
          </article>
        ))}
      </div>
    </section>
  );
}
