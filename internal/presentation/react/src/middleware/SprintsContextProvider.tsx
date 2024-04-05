import React, { useEffect, useState } from "react";
import { Task } from "../domains/Task";
import TaskService from "../services/TaskService";

type SprintsState = {
    [sprint: string]: Boolean;
};

export const SprintsContext = React.createContext({
    getSprints: (): SprintsState => {
        return {} as SprintsState;
    },
    getSelectedSprints: (): Number[] => [],
    toggleSelectedSprint: (sprint: string): void => {},
    setToSprints: (sprint: string): void => {},
    sprintIsSelected: (sprint: string): Boolean => false,
    pullSprintsTrigger: () => {},
});

export default function SprintsProvider({
    children,
}: {
    children: React.ReactNode;
}) {
    const [sprints, setSprints] = useState({} as SprintsState);
    const [trigger, setTrigger] = useState(false);

    useEffect(() => {
        if (trigger) {
            TaskService.sprints().then((s) => {
                setSprints(
                    s.reduce((acc: SprintsState, curr: Number) => {
                        acc[curr.valueOf()] = false;
                        return acc;
                    }, {} as SprintsState)
                );
                setTrigger(false);
            });
        }
    }, [trigger]);

    const getSprints = (): SprintsState => sprints;

    const setToSprints = (sprint: string) => {
        if (sprint && !Object.keys(sprints).includes(sprint)) {
            setSprints({
                ...sprints,
                [sprint]: false,
            });
        }
    };

    const getSelectedSprints = () =>
        Object.entries(sprints)
            .filter(([sprint, selected]) => selected)
            .map(([sprint, selected]) => Number(sprint));

    const toggleSelectedSprint = (sprint: string) => {
        const sprintsCopy = { ...sprints };
        sprintsCopy[sprint] = !Boolean(sprints[sprint]);
        setSprints(sprintsCopy);
    };

    const sprintIsSelected = (sprint: string): Boolean => sprints[sprint];

    const pullSprintsTrigger = () => setTrigger(!trigger);

    return (
        <SprintsContext.Provider
            value={{
                getSprints,
                getSelectedSprints,
                setToSprints,
                toggleSelectedSprint,
                sprintIsSelected,
                pullSprintsTrigger,
            }}
        >
            {children}
        </SprintsContext.Provider>
    );
}
