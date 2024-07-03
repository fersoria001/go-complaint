/* eslint-disable @typescript-eslint/no-explicit-any */

import React from "react";
import { Reply, ReplyReviewType, User } from "../../lib/types";
import { ColorSquare } from "../icons/SquareFillIcon";
import { CommentObject } from "./CommentForm";

/* eslint-disable @typescript-eslint/no-unused-vars */
export type IDOMRect = {
    left: number;
    right: number;
    top: number;
    bottom: number;
    width: number;
    height: number;
}




export const currentDOMRect = (ref: React.RefObject<any> | null): IDOMRect => {
    if (ref && ref.current) {
        const rect = ref.current.getBoundingClientRect();
        return {
            left: rect.left,
            right: rect.right,
            top: rect.top,
            bottom: rect.bottom,
            width: rect.width,
            height: rect.height
        }
    }
    return {
        left: 0,
        right: 0,
        top: 0,
        bottom: 0,
        width: 0,
        height: 0
    }
}

export const difference = (d1: IDOMRect, d2: IDOMRect): IDOMRect => {
    return {
        left: d1.left - d2.left,
        right: d1.right - d2.right,
        top: d1.top - d2.top,
        bottom: d1.bottom - d2.bottom,
        width: d1.width - d2.width,
        height: d1.height - d2.height
    }
}

export const union = (d1: IDOMRect, d2: IDOMRect): IDOMRect => {
    return {
        left: d1.left + d2.left,
        right: d1.right + d2.right,
        top: d1.top + d2.top,
        bottom: d1.bottom + d2.bottom,
        width: d1.width + d2.width,
        height: d1.height + d2.height
    }
}


export type LinePoints = {
    color?: string;
    x1: number;
    x2: number;
    y1: number;
    y2: number;
}
export const xOffset = (pos: IDOMRect) => {
    return pos.left + (pos.width / 2)
}
export const xLimit = (pos: IDOMRect) => {
    return pos.right - (pos.width / 2)
}
export const yCenter = (pos: IDOMRect) => {
    return offsetScrollY(pos.top) - 10
}

const offsetScrollY = (topValue: number) => {
    const scrollTop = window.scrollY || document.documentElement.scrollTop;
    return topValue + scrollTop;
};
const limitScrollY = (bottomValue: number) => {
    const scrollTop = window.scrollY || document.documentElement.scrollTop;
    return bottomValue + scrollTop;
};
export const drawLine = (p1: IDOMRect, p2: IDOMRect): LinePoints => {
    const x1 = p1.left
    const x2 = p2.left - p2.width
    const y1 = offsetScrollY(p1.top)
    const y2 = offsetScrollY(p2.top) - (p2.height / 2)
    return { x1, x2, y1, y2 }
}
export const drawLines = (p1: React.RefObject<any>, childrens: React.RefObject<any>[]): LinePoints[] => {
    const lines: LinePoints[] = []
    const p1Pos = currentDOMRect(p1)
    if (childrens.length > 0) {
        childrens.forEach((child, _) => {
            const children = currentDOMRect(child)
            lines.push(drawLine(p1Pos, children))
        })
    }
    return lines
}



// export const commentPositionSelf = (
//     commentRef: React.RefObject<any>,
//     commentsMap: Map<string, React.RefObject<any>>,
// ): React.RefObject<any> => {
//     let top = 0;
//     let bottom = 0;
//     const arr = Array.from(commentsMap.values())
//     for (let i = 0; i < arr.length; i++) {
//         if (arr[i].current === commentRef.current) {
//             continue
//         }
//         const valuePos = currentDOMRect(commentRef)
//         const lastPos = currentDOMRect(arr[i])
//         const additional = 200
//         const lastPosTopLimit = offsetScrollY(lastPos.top) + additional
//         if (lastPosTopLimit >= offsetScrollY(valuePos.top)) {
//             top = offsetScrollY(lastPos.top) + additional
//             bottom = limitScrollY(lastPos.bottom) + additional
//             // commentRef.current.style.top = `${top}px`
//             // commentRef.current.style.bottom = `${bottom}px`
//         }
//     }

//     return commentRef
// }

export const addToMap = (
    map: Map<string, React.RefObject<any>>,
    commentObject: CommentObject,
    value: React.RefObject<any>,
): Map<string, React.RefObject<any>> => {
    let top = 0;
    let bottom = 0;
    const arr = Array.from(map.values())
    for (let i = 0; i < arr.length; i++) {
        const valuePos = currentDOMRect(value)
        const lastPos = currentDOMRect(arr[i])
        let additional = 200
        if(window.innerWidth < 640) {
            additional = 0
        }

        const lastPosTopLimit = offsetScrollY(lastPos.top) + additional
        if (lastPosTopLimit >= offsetScrollY(valuePos.top)) {
            console.log(window.scrollY , document.documentElement.scrollTop)
            top = offsetScrollY(lastPos.top) + additional
            bottom = limitScrollY(lastPos.bottom) + additional
            value.current.style.top = `${top}px`
            value.current.style.bottom = `${bottom}px`
        }
    }
    const p2 = positionComment(value, Array.from(commentObject.repliesRefs.values()))
    value.current.style.top = `${p2.top}px`
    value.current.style.bottom = `${p2.bottom}px`
    const newMap = new Map([...map, [commentObject.color, value]]);
    return newMap;
}

export const sumChildrens = (A: IDOMRect[]): IDOMRect => {
    let sum: IDOMRect = {
        left: 0,
        right: 0,
        top: 0,
        bottom: 0,
        width: 0,
        height: 0
    }
    for (let i = 0; i < A.length; i++) {
        const top = offsetScrollY(A[i].top)
        let bottom = limitScrollY(A[i].bottom)
        bottom = bottom - A[i].height
        const B: IDOMRect = {
            left: A[i].left,
            right: A[i].right,
            top: top,
            bottom: bottom,
            width: A[i].width,
            height: A[i].height
        }
        sum = union(sum, B)
    }
    return sum
}


export const commentsOverlapY = (
    value: IDOMRect,
    arr: React.RefObject<any>[],
): IDOMRect => {
    let top = value.top;
    let bottom = value.bottom;
    for (let i = 0; i < arr.length; i++) {
        const lastPos = currentDOMRect(arr[i])
        const additional = lastPos.height
        //const lastPosTopLimit = offsetScrollY(lastPos.top) + additional
        if (value.top <= offsetScrollY(lastPos.bottom)) {
            top = top + additional
            bottom = bottom + additional
        }
    }
    return {
        left: 0,
        right: 0,
        top: top,
        bottom: bottom,
        width: 0,
        height: 0
    }
}

/**
 * 
 * @param p1 
 * @param childrens 
 * @returns 
 * @throws {Error} if childrens are empty
 */
export const positionComment = (
    p1: React.RefObject<any>,
    childrens2: React.RefObject<any>[],
    commentsMap?: Map<string, React.RefObject<any>>,
    childrens1?: React.RefObject<any>[],
): IDOMRect => {
    const p1Pos = currentDOMRect(p1)
    if (childrens2?.length === 0 || !childrens2) {
        throw new Error("Childrens are empty")
    }
    const childrens2Pos = childrens2.map((child) => currentDOMRect(child))
    const childrens2Sum = sumChildrens(childrens2Pos)
    let childrensSum = childrens2Sum
    if (childrens1) {
        const childrens1Pos = childrens1.map((child) => currentDOMRect(child))
        const childrens1Sum = sumChildrens(childrens1Pos)
        childrensSum = difference(childrens2Sum, childrens1Sum)
    }
    const height = p1Pos.height
    const i = childrens2Pos.length + 1
    const top = childrensSum.top + offsetScrollY(p1Pos.top)
    const bottom = childrensSum.bottom + limitScrollY(p1Pos.top)
    const newTop = top / i
    const newBottom = bottom / i + height
    let result = {
        left: p1Pos.left,
        right: p1Pos.right,
        top: newTop,
        bottom: newBottom,
        width: p1Pos.width,
        height: p1Pos.height
    }
    if (commentsMap) {
        const commentsArray = Array.from(commentsMap.values())
        const filtered = commentsArray.filter((comment) => comment.current !== p1.current)
        result = commentsOverlapY(result, filtered)
    }
    return result
}


/**
 * 
 * @param commentsMap
 * @param repliesMap
 * @param colorKey
 * @return new commentsMap
 * @throws {Error} if comment not found in commentsMap
 */
export const commentHeight = (
    commentsMap: Map<string, React.RefObject<any>>,
    repliesMap1: Map<string, React.RefObject<any>>,
    repliesMap2: Map<string, React.RefObject<any>>,
    colorKey: string,
): Map<string, React.RefObject<any>> => {
    const comment = commentsMap.get(colorKey)
    if (!comment) {
        throw new Error(`Comment ${colorKey} not found in commentsMap`)
    }
    const childrens1 = Array.from(repliesMap1.values())
    const childrens2 = Array.from(repliesMap2.values())
    const newPosition = positionComment(comment, childrens1, commentsMap, childrens2)
    comment.current.style.top = `${newPosition.top}px`
    comment.current.style.bottom = `${newPosition.bottom}px`
    const newMap = new Map(commentsMap.entries())
    newMap.delete(colorKey)
    newMap.set(colorKey, comment)
    return newMap
}
export const generateSquares = (): ColorSquare[] => {
    return [
        {
            selected: false,
            picked: false,
            fixed: false,
            color: "#ef4444"
        },
        {
            selected: false,
            picked: false,
            fixed: false,
            color: "#eab308"
        },
        {
            selected: false,
            picked: false,
            fixed: false,
            color: "#3b82f6"
        },
    ]
}



/**
 * 
 * @param colorSquares 
 * @param color 
 * @returns 
 * @throws {Error} if color not found in colorSquares
 */
export const unpickColor = (
    colorSquares: ColorSquare[],
    color: string,
): ColorSquare[] => {
    const squareIndex = colorSquares.findIndex((square) => square.color === color);
    if (squareIndex !== -1) {
        const newSquares = colorSquares.map((square, idx) => idx === squareIndex ? {
            ...square,
            selected: false,
            picked: false
        } : square);
        return newSquares;
    }
    throw new Error(`Color ${color} not found in colorSquares`);
}


export const colorAlreadyPicked = (colorSquares: ColorSquare[], color: string): boolean => {
    const squareIndex = colorSquares.findIndex((square) => square.color === color);
    if (squareIndex !== -1) {
        const square = colorSquares[squareIndex];
        if (square.picked) {
            return true
        }
    }
    return false
}

/**
 * 
 * @param colorSquares 
 * @param color 
 * @returns 
 * @throws {Error} if color not found in colorSquares
 * @throws {ColorAlreadyPickedError} if color already picked
 */
export const colorPicked = (
    colorSquares: ColorSquare[],
    color: string,
): ColorSquare[] => {
    const squareIndex = colorSquares.findIndex((square) => square.color === color);
    if (squareIndex !== -1) {
        const square = colorSquares[squareIndex];
        if (square.picked) {
            throw new ColorAlreadyPickedError(`${color} color already picked`);
        }
        const newSquares = colorSquares.map((square, idx) => idx === squareIndex ? {
            ...square,
            selected: false,
            picked: true
        } : square);
        return newSquares;
    }
    throw new Error(`Color ${color} not found in colorSquares`);
}

export const loadColorSquares = (
    colorSquares: ColorSquare[],
    colorsPicked: string[]): ColorSquare[] => {
    let newSquares: ColorSquare[] = colorSquares;
    for (let i = 0; i < colorSquares.length; i++) {
        if (colorAlreadyPicked(colorSquares, colorSquares[i].color)) {
            newSquares = unpickColor(newSquares, colorSquares[i].color);
        }
    }
    for (let i = 0; i < colorsPicked.length; i++) {
        newSquares = colorPicked(newSquares, colorsPicked[i]);
    }
    newSquares = fixSquaresPos(newSquares, false);
    return newSquares;
}

export const loadReplyReviews = (
    replyReviews: ReplyReviewType[],
    repliesMap: Map<string, React.RefObject<HTMLLIElement>>,
): { loadedComments: CommentObject[], colorsPicked: string[] } => {
    const loadedComments: CommentObject[] = [];
    const colorsPicked: string[] = [];
    replyReviews.sort((a, b) => parseInt(a.createdAt) - parseInt(b.createdAt));
    for (let i = 0; i < replyReviews.length; i++) {
        const replyReview = replyReviews[i];
        let comment = createComment(replyReview.color, replyReview.reviewer);
        if (replyReview.review.comment != "") {
            comment = setBody(comment, replyReview.review.comment);
            comment = setDone(comment, true);
        }
        for (let j = 0; j < replyReview.replies.length; j++) {
            const replyRef = repliesMap.get(replyReview.replies[j].id);
            if (!replyRef) {
                throw new Error(`replyRef ${replyReview.replies[j].id} not found in repliesMap`);
            }

            if (!replyAlreadyExists(comment, replyReview.replies[j])) {
                comment = addReply(comment, replyReview.replies[j], replyRef);
            }
        }
        loadedComments.push(comment);
        colorsPicked.push(replyReview.color);
    }
    return { loadedComments, colorsPicked };
}


/**
 * 
 * @param replyReviews 
 * @param colorSquares 
 * @param comments 
 * @param repliesMap 
 * @returns 
 * @throws {Error} if replyRef not found in repliesMap
 * @throws {Error} if addReply throws an error because reply already exists in comment
 * @throws {Error} if addComment throws an error because comment already exists in comments
 * @throws {ColorAlreadyPickedError} if colorPicked throws an error because color already picked
 * @throws {Error} if colorPicked throws an error because color not found in colorSquares
 */
export const load = (
    user: User,
    replyReviews: ReplyReviewType[],
    colorSquares: ColorSquare[],
    repliesMap: Map<string, React.RefObject<HTMLLIElement>>,
): { newComments: CommentObject[], newSquares: ColorSquare[], actualColor: string } => {
    let actualColor = "";
    if (!replyReviews) {
        const newSquares = loadColorSquares(colorSquares, []);
        return { newComments: [], newSquares, actualColor };
    }
    const { loadedComments, colorsPicked } = loadReplyReviews(replyReviews, repliesMap);
    let newSquares = loadColorSquares(colorSquares, colorsPicked);
    const unDoneComment = loadedComments.find((comment) => comment.reviewer.email === user.email && !comment.done);
    if (unDoneComment) {
        newSquares = newSquares.map((square) => square.color === unDoneComment.color ? { ...square, selected: true, fixed: true } : { ...square, fixed: true });
        actualColor = unDoneComment.color;
    }
    return { newComments: loadedComments, newSquares: newSquares, actualColor };
}

/**
 * 
 * @param comments 
 * @param color 
 * @param body 
 * @returns newComments
 * @throws {Error} if comment does not exist in comments
 */
export const markCommentAsDone = (
    comments: CommentObject[],
    color: string,
    body: string,
): CommentObject[] => {
    let commentObject = comments.find((comment) => comment.color == color);
    if (!commentObject) {
        throw new Error(`commentObject not found for color ${color}`);
    }
    commentObject = setBody(commentObject, body);
    commentObject = setDone(commentObject, true);
    const newComments = comments.map((comment) => comment.color == color ? commentObject : comment);
    return newComments;
}

/**
 * 
 * @param comments 
 * @param colorSquares 
 * @param color 
 * @returns 
 * @throws {Error} if comment does not exist in comments
 * @throws {Error} if comment has 0 replies
 */
export const markCommentAsEditable = (
    comments: CommentObject[],
    colorSquares: ColorSquare[],
    color: string,
): { newComments: CommentObject[], newSquares: ColorSquare[] } => {
    const comment = comments.find((comment) => comment.color === color);
    if (!comment) {
        throw new Error(`Comment ${color} not found in comments`);
    }
    if (comment.selectedReplies.length < 0) {
        throw new Error(`Comment ${color} has 0 replies`);
    }
    let newColors = unpickColor(colorSquares, color);
    newColors = selectAColor(color, newColors);
    const newComments = comments.map((c) => c.color === color ? setDone(c, false) : c);
    return { newComments: newComments, newSquares: newColors };
}

export const removeFromMap = (
    map: Map<string, React.RefObject<any>>,
    key: string
): Map<string, React.RefObject<any>> => {
    const newMap = new Map(map.entries());
    newMap.delete(key);
    return newMap;
}

/**
 * @param comments
 * @param color
 * @returns newComments
 * 
 */
export const commentAlreadyExists = (
    comments: CommentObject[],
    comment: CommentObject,
): boolean => {
    const exists = comments.find((c) => c.color === comment.color);
    return exists ? true : false;
}


/**
 * Adds a comment to a new comments array 
 * with previous object + new object and returns the new array
 * @param comments 
 * @param comment 
 * @returns newComments
 * @throws {Error} if comment already exists in comments
 */
export const addComment = (
    comments: CommentObject[],
    comment: CommentObject,
): CommentObject[] => {
    const index = comments.findIndex((c) => c.color === comment.color);
    if (index !== -1) {
        throw new Error(`Comment ${comment.color} already exists in comments`);
    }
    const newComments = [...comments, comment];
    return newComments;
}

export const swapComment = (
    comments: CommentObject[],
    comment: CommentObject,
    index: number
): CommentObject[] => {
    const newComments = comments.map((c, i) => i === index ? comment : c);
    return newComments;
}

/**
 * 
 * @param comments 
 * @param commentsMap 
 * @param color 
 * @returns 
 * @throws {Error} if comment does not exist in comments
 */
export const removeComment = (
    comments: CommentObject[],
    commentsMap: Map<string, React.RefObject<any>>,
    color: string,
): { objects: CommentObject[], refs: Map<string, React.RefObject<any>> } => {
    const comment = comments.find((comment) => comment.color === color);
    if (!comment) {
        throw new Error(`Comment ${color} not found in comments`);
    }
    comment.repliesRefs.forEach((ref) => {
        if (ref) {
            ref.current!.style.borderColor = "";
            ref.current!.style.borderWidth = "0px";
        }
    });
    const newComments = comments.filter((comment) => comment.color !== color);
    const newMap = removeFromMap(commentsMap, color);
    return { objects: newComments, refs: newMap };
}

/**
 * 
 * @param comment 
 * @param reply 
 * @param replyRef 
 * @returns 
 * @throws {Error} if reply does not exist in comment
 */
export const removeReply = (
    replyRef: React.RefObject<HTMLLIElement> | null
): void => {
    replyRef!.current!.style.borderColor = "";
    replyRef!.current!.style.borderWidth = "0px";
    replyRef!.current!.style.borderRadius = "0px";
}



export const replyAlreadyExists = (
    comment: CommentObject,
    reply: Reply
): boolean => {
    const exists = comment.selectedReplies.find((r) => r.id === reply.id);
    return exists ? true : false;
}
/**
 * Colors the reply ref, ads the reply to the comment and 
 * updates the reply refs map of the comment
 * @param comment 
 * @param reply 
 * @param replyRef 
 * @returns  newCommentObject[]
 * @throws {Error} if reply already exists in comment
 */
export const addReply = (
    comment: CommentObject,
    reply: Reply,
    replyRef: React.RefObject<HTMLLIElement>
): CommentObject => {
    const selectedReplies = comment.selectedReplies;
    const exists = replyAlreadyExists(comment, reply);
    if (exists) {
        throw new Error(`Reply ${reply.id} already exists in comment ${comment.color}`);
    }
    replyRef!.current!.style.borderColor = comment.color;
    replyRef!.current!.style.borderWidth = "1px";
    replyRef!.current!.style.borderRadius = "5px";
    selectedReplies.push(reply);
    const refMap = new Map(comment.repliesRefs.entries());
    refMap.set(reply.id, replyRef);
    return { ...comment, selectedReplies: selectedReplies, repliesRefs: refMap };
}
export const createComment = (
    color: string,
    reviewer: User,
): CommentObject => {
    return {
        reviewer: reviewer,
        color: color,
        comment: "",
        done: false,
        repliesRefs: new Map(),
        selectedReplies: [],
    };
}
export const setReviewer = (
    comment: CommentObject,
    reviewer: User
): CommentObject => {
    return { ...comment, reviewer: reviewer }
}
export const setBody = (
    comment: CommentObject,
    body: string
): CommentObject => {
    return { ...comment, comment: body }
}
export const setDone = (
    comment: CommentObject,
    done: boolean
): CommentObject => {
    return { ...comment, done: done }
}


/**
 * 
 * @param selectedColor 
 * @param colorSquares 
 * @returns {ColorSquare[]}
 * @throws {Error} if color not found in colorSquares
 * @throws {ColorAlreadyPickedError} if color already picked
 */
export const selectAColor = (
    selectedColor: string,
    colorSquares: ColorSquare[],
): ColorSquare[] => {
    const square = colorSquares.find((square) => square.color === selectedColor);
    if (!square) {
        throw new Error(`Color ${selectedColor} not found in colorSquares`);
    }
    if (square.fixed) {
        throw new ColorsFixedError(`finish your review by adding a comment first or delete it`);
    }
    if (square.picked) {
        throw new ColorAlreadyPickedError(`${selectedColor} color already picked`);
    }
    const newSquare = { ...square, selected: true };
    const newSquares = colorSquares.map((v) => v.color === selectedColor ? newSquare : { ...v, selected: false });
    return newSquares;
}

export const fixSquaresPos = (
    colorSquares: ColorSquare[],
    value = true
): ColorSquare[] => {
    const newSquares = colorSquares.map((square) => {
        return { ...square, fixed: value }
    });
    return newSquares;
}



export type FeedbackErrorType = { [key: string]: { [key: string]: string } }

export type ColorError = {
    colorKey: string;
    error: string;
}

export class ColorAlreadyPickedError extends Error {
    constructor(message: string) {
        super(message);
        this.name = "ColorAlreadyPickedError";
    }
}

export class ColorsFixedError extends Error {
    constructor(message: string) {
        super(message);
        this.name = "ColorsFixedError";
    }
}

