import React from "react";
import Button from "./Button";
import Task from "./Task";

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
      <h1 className="text-3xl font-bold text-stone-600 mb-5">Tasks</h1>
      <div className="w-1/2 flex justify-between h-8 mb-6">
        <input className="bg-stone-300 w-4/5 p-3"></input>
        <Button title={"Add Task"} style={"text-lg font-bold text-stone-600"} />
      </div>
      <div className="w-1/2 flex flex-col space-y-4 bg-stone-200 p-3">
        <Task task={"Practice, practice, practice !"} />
        <Task task={"Learn advance topic"} />
      </div>
    </div>
  );
}