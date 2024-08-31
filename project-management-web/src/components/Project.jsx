import React from "react";
import Button from "./Button";
import Task from "./Task";
import Input from "./Input";

export default function Project({ activeProject }) {
  return (
    <div className="px-10 py-16 w-full">
      <h1 className="text-3xl font-bold text-stone-600 mb-5">
        {activeProject.title}
      </h1>
      <p className="text-stone-400 mb-4 font-bold text-lg">
        {activeProject.date}
      </p>
      <p className="text-stone-600 mb-4 font-bold text-lg">
        {activeProject.description1}
      </p>
      <p className="text-stone-600 mb-4 font-bold text-lg">
        {activeProject.description2}
      </p>
      <div className="border-t border-gray-300 my-4"></div>
      <h1 className="text-3xl font-bold text-stone-600 mb-5">Tasks</h1>
      <div className="w-1/2 flex justify-between h-8 mb-6">
        <Input style="bg-stone-300 w-4/5 p-3" />
        <Button title={"Add Task"} style={"text-lg font-bold text-stone-600"} />
      </div>
      <div className="w-1/2 flex flex-col space-y-4 bg-stone-200 p-3">
        <Task task={"Practice, practice, practice !"} />
        <Task task={"Learn advance topic"} />
      </div>
    </div>
  );
}
