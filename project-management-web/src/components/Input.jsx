import React from "react";

export default function Input({ handleChange, style, name }) {
  return <input className={style} onChange={handleChange} name={name}></input>;
}
