package queries

type ComplaintQuery struct {
	Term         string `json:"term"`
	ID           string `json:"id"`
	ReceiverID   string `json:"receiver_id"`
	ReceiverName string `json:"receiver_name"`
	AuthorID     string `json:"author_id"`
	Limit        int    `json:"limit"`
	Offset       int    `json:"offset"`
	UserID       string `json:"user_id"`
	Status       string `json:"status"`
	AfterDate    string `json:"after_date"`
	BeforeDate   string `json:"before_date"`
}

// func (query ComplaintQuery) History(
// 	ctx context.Context,
// ) (dto.ComplaintListDTO, error) {
// 	if query.ReceiverID == "" {
// 		return dto.ComplaintListDTO{}, ErrBadRequest
// 	}
// 	received, err := repositories.MapperRegistryInstance().Get(
// 		"Complaint",
// 	).(repositories.ComplaintRepository).FindAll(
// 		ctx,
// 		find_all_complaints.ByReceiverAndStatus(
// 			query.ReceiverID,
// 			complaint.IN_HISTORY.String(),
// 		),
// 	)
// 	if err != nil {
// 		return dto.ComplaintListDTO{}, err
// 	}
// 	if query.AfterDate != "" && query.BeforeDate != "" {
// 		after, err := common.NewDateFromString(query.AfterDate)
// 		if err != nil {
// 			return dto.ComplaintListDTO{}, err
// 		}
// 		before, err := common.NewDateFromString(query.BeforeDate)
// 		if err != nil {
// 			return dto.ComplaintListDTO{}, err
// 		}
// 		for c := range received.Iter() {
// 			if !c.CreatedAt().Date().After(after.Date()) && !c.CreatedAt().Date().Before(before.Date()) {
// 				received.Remove(c)
// 			}
// 		}
// 	}

// 	if query.Term != "" {
// 		trie := trie.NewTrie()
// 		for c := range received.Iter() {
// 			trie.InsertText(c.ID().String(), c.AuthorFullName(), " ")
// 			trie.InsertText(c.ID().String(), c.Message().Title(), " ")
// 			trie.InsertText(c.ID().String(), c.Message().Description(), " ")
// 			trie.InsertText(c.ID().String(), c.Message().Body(), " ")
// 			trie.InsertText(c.ID().String(), c.Status().String(), "_")
// 		}
// 		ids := trie.Search(query.Term)
// 		if ids == nil {
// 			return dto.ComplaintListDTO{
// 				Complaints:    []dto.ComplaintDTO{},
// 				Count:         0,
// 				CurrentLimit:  query.Limit,
// 				CurrentOffset: query.Offset,
// 			}, nil
// 		}
// 		slice := received.ToSlice()
// 		for c := range slice {
// 			if !ids.Contains(slice[c].ID().String()) {
// 				received.Remove(slice[c])
// 			}
// 		}
// 	}

// 	count := received.Cardinality()
// 	complaints := make([]dto.ComplaintDTO, 0, count)

// 	receivedSlice := received.ToSlice()

// 	slices.SortStableFunc(receivedSlice, func(i, j complaint.Complaint) int {
// 		if i.CreatedAt().Date().After(j.CreatedAt().Date()) {
// 			return -1
// 		}
// 		if i.CreatedAt().Date().Before(j.CreatedAt().Date()) {
// 			return 1
// 		}
// 		return 0
// 	})

// 	for _, c := range receivedSlice {
// 		complaints = append(complaints, dto.NewComplaintDTO(c))
// 	}
// 	offset := query.Offset
// 	limit := query.Limit
// 	length := len(complaints)
// 	offsetLimit := offset + limit
// 	//offset: 0 | < len | > len
// 	//limit: 10 | < len | > len
// 	if offset > length {
// 		return dto.ComplaintListDTO{}, fmt.Errorf("offset is greater than the length of the list")
// 	}
// 	if offset+limit > length {
// 		offsetLimit = offset + (length - offset)
// 	}
// 	complaints = complaints[offset:offsetLimit]
// 	return dto.ComplaintListDTO{
// 		Complaints:    complaints,
// 		Count:         count,
// 		CurrentLimit:  query.Limit,
// 		CurrentOffset: query.Offset,
// 	}, nil
// }
