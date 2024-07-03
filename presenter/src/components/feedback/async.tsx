import { AddCommentMutation, AddReplyMutation, CreateFeedbackMutation, DeleteCommentMutation, EndFeedbackMutation, Mutation, RemoveReplyMutation } from "../../lib/mutations";
import { CreateAFeedback, Reply } from "../../lib/types";
import { CommentObject } from "./CommentForm";

export async function saveNewComment(
    comment: CommentObject,
    complaintID: string,
    enterpriseID: string,
): Promise<boolean> {
    return Mutation<CreateAFeedback>(
        CreateFeedbackMutation,
        {
            enterpriseID,
            complaintID,
            color: comment.color,
        }
    ).then((res) => {
        return res;
    });
}

export async function saveNewReply(
    reply: Reply,
    comment: CommentObject,
    enterpriseID: string,
    feedbackID: string,
    reviewerID: string,
): Promise<boolean> {
    return Mutation<CreateAFeedback>(
        AddReplyMutation,
        {
            enterpriseID,
            feedbackID,
            reviewerID,
            color: comment.color,
            repliesID: [reply.id],
        }
    ).then((res) => {
        return res;
    }).catch((err) => {
        console.error(err);
        return false;
    });
}

export async function removeNewReply(
    reply: Reply,
    comment: CommentObject,
    enterpriseID: string,
    feedbackID: string,
): Promise<boolean> {
    return Mutation<CreateAFeedback>(
        RemoveReplyMutation,
        {
            enterpriseID,
            feedbackID,
            color: comment.color,
            repliesID: [reply.id],
        }
    ).then((res) => {
        return res;
    }).catch((err) => {
        console.error("removeNewReplyMutation", err);
        return false;
    });
}
export async function setCommentBody(
    comment: CommentObject,
    enterpriseID: string,
    feedbackID: string,
): Promise<boolean> {
    return Mutation<CreateAFeedback>(
        AddCommentMutation,
        {
            enterpriseID,
            feedbackID,
            color: comment.color,
            comment: comment.comment,
        }
    ).then((res) => {
        return res;
    });
}

export async function deleteComment(
    comment: CommentObject,
    enterpriseID: string,
    feedbackID: string,
): Promise<boolean> {
    return Mutation<CreateAFeedback>(
        DeleteCommentMutation,
        {
            enterpriseID,
            feedbackID,
            color: comment.color,
        }
    ).then((res) => {
        return res;
    });
}


export async function submitFeedback(
    enterpriseID: string,
    feedbackID: string,
): Promise<boolean> {
    return Mutation<CreateAFeedback>(
        EndFeedbackMutation,
        {
            enterpriseID,
            feedbackID,
        }
    ).then((res) => {
        return res;
    })
}
