import React, { useState, useEffect } from "react";
import { useRouter } from "next/router";
import style from "./Tab.module.css";
import slugify from "../utils/slugify";

interface Props {
    initialTab: {
        tab?: string;
    };
    tab?: string;
}
const Tabs: React.FC<Props> = ({children, initialTab}) => {
    let firstTab: string;
    //if(React.Children.count(children) > 0) {
    let initalTabElement = React.Children.toArray(children)[0] as React.ReactElement<Props>;
    firstTab = initalTabElement.props.tab as string;
    //}
    const [currentTab, setCurrentTab] = useState(slugify(firstTab as string));
    const router = useRouter();


    const handleTabClick = (e: React.MouseEvent<HTMLAnchorElement, MouseEvent>, tab: string) => {
        e.preventDefault();
        console.log(tab)
        setCurrentTab(slugify(tab));
    }

    useEffect(() => {
        if(initialTab.tab) {
            setCurrentTab(initialTab.tab);
            console.log(initialTab.tab)
        }
    }, []);

    useEffect(() => {
        router.push(`${router.pathname}?tab=${slugify(currentTab)}`, undefined, { shallow: true });
    }, [currentTab]);

    return (
        <>
        <ul className={style.tabs}>
            {React.Children.map(children, (child) => {
                let item = child as React.ReactElement<Props>;
                return (
                    <a href="#" onClick={(e) => handleTabClick(e, item.props.tab as string)}>
                        <li key={item.props.tab} className={slugify(item.props.tab as string) === currentTab ? style.active__tab : style.tab}>
                            <span>{item.props.tab}</span>
                        </li>
                    </a>);
            })}
        </ul>

        {React.Children.map(children, (child) => {
            let item = child as React.ReactElement<Props>;
            if (slugify(item.props.tab as string) === currentTab) {
                return child;
            }
        })}

        </>
    );
}

export default Tabs;