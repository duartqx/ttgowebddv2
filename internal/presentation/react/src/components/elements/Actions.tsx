import { useContext, useState } from "react";
import FilterTasksForm from "../tasks/FilterTasksForm";
import { TaskCreate, TaskFilter } from "../../domains/Task";
import Tab from "./Tab";
import CreateTaskForm from "../tasks/CreateTaskForm";
import ActionsButton from "./ActionsButton";
import ActionsTabs from "./ActionsTabs";
import ActionsFormContainer from "./ActionsFormContainer";
import Backdrop from "./Backdrop";
import { AuthContext } from "../../middleware/AuthContextProvider";
import ArrowRightEnd from "../../icons/ArrowRightEndIcon";
import { useNavigate } from "react-router-dom";

type ActionsFormsProps = {
    setTaskFilter: (tf: TaskFilter) => void;
    newTaskHandler: (t: TaskCreate) => void;
};

type Tab = "Filter" | "Create";

export default function Actions({
    setTaskFilter,
    newTaskHandler,
}: ActionsFormsProps) {
    const { logout } = useContext(AuthContext);
    const [formsVisible, setFormsVisibile] = useState(false);
    const [selectedTab, setSelectedTab] = useState("Filter" as Tab);
    const navigate = useNavigate();

    const isSelected = (t: Tab) => () => selectedTab === t;
    const selHandler = (t: Tab) => () => setSelectedTab(t);

    return (
        <Backdrop
            isOpen={formsVisible}
            children={
                <>
                    <nav
                        className="
                        flex bg-zinc-950 h-[60px]
                        justify-between items-center
                        fixed left-0 top-0 w-[100vw]
                        "
                    >
                        <ActionsButton
                            active={formsVisible}
                            setActive={() => setFormsVisibile(!formsVisible)}
                        />
                        <button
                            className="
                                h-3/4 flex justify-center mr-8 bg-transparent
                                hover:border-red-500 hover:text-red-500
                                transform-all duration-500 ease-in-out
                            "
                            onClick={() => {
                                logout();
                                navigate("/");
                            }}
                        >
                            <ArrowRightEnd />
                        </button>
                    </nav>
                    <div
                        className={`
                        ${
                            !formsVisible
                                ? "opacity-0 -top-[100vh]"
                                : "opacity-100 top-[4.5rem]"
                        }
                        fixed bg-zinc-800 w-[400px] rounded-md
                        left-1/2 -translate-x-[200px]
                        shadow-lg shadow-zinc-950 flex flex-col
                        transform-all duration-300 ease-in-out min-h-[565px]
                    `}
                    >
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

                        {/** Filter and Create Forms */}
                        <ActionsFormContainer
                            isHidden={selectedTab !== "Filter"}
                            children={
                                <FilterTasksForm
                                    setTaskFilter={setTaskFilter}
                                />
                            }
                        />
                        <ActionsFormContainer
                            isHidden={selectedTab !== "Create"}
                            children={
                                <CreateTaskForm
                                    setNewTaskHandler={newTaskHandler}
                                />
                            }
                        />
                        {/** End Forms */}
                    </div>
                </>
            }
        />
    );
}
