import React from "react";

export default function Project({ activeProject }) {
  return (
    <div className="px-10 py-16 w-full">
      <h1 className="text-3xl font-bold text-stone-600 mb-5">Learning React</h1>
      <p className="text-stone-400 mb-4 font-bold text-lg">Dec 29, 2024</p>
      <p className="text-stone-600 mb-4 font-bold text-lg">
        Learn react from the group up.
      </p>
      <p className="text-stone-600 mb-4 font-bold text-lg">
        Start with the basics, finish with advanced knowledge.
      </p>
      <div class="border-t border-gray-300 my-4"></div>
    </div>
  );
}
