package lib

type DicomAPIRequest struct {
	PatientName      string
	flag             int
	StudyDescription string
	StudyDate        string
	StudyId          string
	List             []Series
}

type Series struct {
	seriesUid         string
	seriesNumber      int
	seriesTime        string
	seriesDescription string
	seriesDate        string
	InstanceList      []Instance
}

type Instance struct {
	ImageId             string
	FrameOfReferenceUID string
}
