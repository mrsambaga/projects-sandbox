import React from "react";

export default function Input({ handleChange, style, name, value }) {
  return (
    <input
      className={style}
      onChange={handleChange}
      name={name}
      value={value}
    ></input>
  );
}
