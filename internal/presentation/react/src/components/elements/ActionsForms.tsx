import { useState } from "react";
import FilterTasksForm from "../tasks/FilterTasksForm";
import { TaskCreate, TaskFilter } from "../../domains/Task";
import Tab from "./Tab";
import CreateTaskForm from "../tasks/CreateTaskForm";
import ActionsButton from "./ActionsButton";
import ActionsTabs from "./ActionsTabs";
import ActionsFormContainer from "./ActionsFormContainer";
import Backdrop from "./Backdrop";

type ActionsFormsProps = {
    setTaskFilter: (tf: TaskFilter) => void;
    newTaskHandler: (t: TaskCreate) => void;
};

type Tab = "Filter" | "Create";

export default function ActionsForms({
    setTaskFilter,
    newTaskHandler,
}: ActionsFormsProps) {
    const [formsVisible, setFormsVisibile] = useState(false);
    const [selectedTab, setSelectedTab] = useState("Filter" as Tab);

    const isSelected = (t: Tab) => () => selectedTab === t;
    const selHandler = (t: Tab) => () => setSelectedTab(t);

    return (
        <Backdrop
            isOpen={formsVisible}
            children={
                <>
                    <div
                        className={`
                        ${
                            !formsVisible
                                ? "opacity-0 bottom-14"
                                : "opacity-100 bottom-20"
                        }
                        fixed bg-neutral-900 w-[400px] rounded-md
                        left-1/2 -translate-x-[200px]
                        shadow-lg shadow-neutral-950 flex flex-col
                        transform-all duration-200 ease-in-out min-h-[585px]
                    `}
                    >
                        {/** Filter and Create Forms */}
                        <ActionsFormContainer
                            label="Filter"
                            isHidden={selectedTab !== "Filter"}
                            children={
                                <FilterTasksForm
                                    setTaskFilter={setTaskFilter}
                                />
                            }
                        />
                        <ActionsFormContainer
                            label="Create"
                            isHidden={selectedTab !== "Create"}
                            children={
                                <CreateTaskForm
                                    setNewTaskHandler={newTaskHandler}
                                />
                            }
                        />
                        {/** End Forms */}

                        <ActionsTabs
                            which={selectedTab === "Filter"}
                            children={
                                <>
                                    <Tab
                                        label="Filter"
                                        isSelected={isSelected("Filter")}
                                        onClickHandler={selHandler("Filter")}
                                    />
                                    <Tab
                                        label="Create"
                                        isSelected={isSelected("Create")}
                                        onClickHandler={selHandler("Create")}
                                    />
                                </>
                            }
                        />
                        {/** End ActionsTabs */}
                    </div>

                    <ActionsButton
                        active={formsVisible}
                        setActive={() => setFormsVisibile(!formsVisible)}
                    />
                </>
            }
        />
    );
}
