import React from "react";
import Button from "./Button";

export default function Task({ task, handleClear }) {
  return (
    <div className="flex justify-between">
      <p className="text-lg font-bold text-stone-600">{task}</p>
      <Button title={"clear"} style={"text-lg font-bold text-stone-600"} />
    </div>
  );
}
