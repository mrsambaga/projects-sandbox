import React from "react";
import Input from "./Input";

export default function Form({ title, onChange, field }) {
  return (
    <div>
      <h2 className="text-stone-600 font-bold text-lg">{title}</h2>
      <Input
        style="bg-stone-300 p-3 w-full"
        handleChange={onChange}
        name={field}
      />
    </div>
  );
}
