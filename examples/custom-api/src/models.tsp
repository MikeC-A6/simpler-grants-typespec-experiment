import "@common-grants/core";
import "@typespec/json-schema";
import "@typespec/http";

using JsonSchema;
using CommonGrants;

@jsonSchema
namespace CustomAPI.CustomModels;

@doc("A funding opportunity number")
model OpportunityNumber extends Models.CustomField {
    name: "Funding Opportunity Number";
    type: Models.CustomFieldType.string;

    @example("ABC-2025-001")
    value: string;

    description: "The unique identifier for a given opportunity within this API";
}

@doc("Metadata about an opportunity")
model OpportunityMetadata {
    @doc("When this opportunity was last modified")
    lastModified: string;
    
    @doc("The user who last modified this opportunity")
    modifiedBy: string;
    
    @doc("The status of this opportunity")
    status: "draft" | "published" | "closed";
    
    @doc("The deadline for applications")
    applicationDeadline: string;

    @doc("The review status of this opportunity")
    reviewStatus: "pending" | "in_review" | "reviewed" | "approved" | "rejected";

    @doc("Categories of funding this opportunity belongs to")
    fundingCategories: string[];
}

@doc("A custom opportunity model extending the base opportunity")
model CustomOpportunity {
    @doc("The base opportunity fields")
    id: Models.uuid;
    source: string;
    title: string;
    description: string;
    fundingDetails: Models.FundingDetails;
    applicationTimeline?: Models.Event[];

    @doc("Metadata about this opportunity")
    metadata: OpportunityMetadata;

    @doc("Custom fields for this opportunity")
    customFields: {
        @doc("The funding opportunity number")
        fundingOpportunityNumber: OpportunityNumber;
    };
}
