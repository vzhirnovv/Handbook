import React from "react";
import ReactDOM from "react-dom";

const Modal = props => {
    const { open, onClose, children } = props;
    return <>{open && ReactDOM.createPortal({})}</>;
};
