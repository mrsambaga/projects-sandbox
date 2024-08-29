import React from "react";
import ProjectItem from "./ProjectItem";
import Button from "./Button";

export default function SideBar() {
  return (
    <div className="w-1/3 h-full px-8 py-16 bg-stone-900 text-stone-50 md:w-72 rounded-r-xl flex-col space-y-8">
      <h1 className="font-bold uppercase md:text-xl text-stone-200">
        YOUR PROJECTS
      </h1>
      <Button
        title={"+ Add Project"}
        style="px-6 py-2 rounded-md bg-stone-800 text-stone-50 hover:bg-stone-950"
      />
      <div>
        <Button
          title={"Project 1"}
          style="w-full text-left px-2 py-1 rounded-sm my-1 hover:text-stone-200 hover:bg-stone-800"
        />
        <Button
          title={"Project 2"}
          style="w-full text-left px-2 py-1 rounded-sm my-1 hover:text-stone-200 hover:bg-stone-800"
        />
      </div>
    </div>
  );
}
