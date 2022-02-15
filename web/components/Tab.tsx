import React from "react";
interface Props {
    tab: string;
}

const Tab: React.FC<Props> = ({children}) => {
    return (
        <>
        {React.Children.map(children, (child) => 
            child
        )}
        </>
    );
}

export default Tab;