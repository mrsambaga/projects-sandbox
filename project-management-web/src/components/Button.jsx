import React from "react";

export default function Button({ title, clickHandler, style }) {
  return (
    <button className={style} onClick={clickHandler}>
      {title}
    </button>
  );
}
