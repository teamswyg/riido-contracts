import { jsonText } from "../contractModel.js";

export default function JsonBlock({ title, value }) {
  return (
    <section className="json-block">
      <h3>{title}</h3>
      <pre>{jsonText(value)}</pre>
    </section>
  );
}
