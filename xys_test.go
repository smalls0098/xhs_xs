package xs

import "testing"

func Test_xys(t *testing.T) {
	params := `/api/sns/web/v1/feed{"source_note_id":"6864faa6000000000b01f72a","image_formats":["jpg","webp","avif"],"extra":{"need_body_topic":"1"},"xsec_source":"pc_feed","xsec_token":"AB7TxBzN4BsJ-j5Vezf6B15-a8APH66lh4_CZJq6zC3a0="}`
	a1 := "19812c441384uqq7gm813oi08vfpk7d1zt2jbxyil30000150545"
	t.Log(XYS(params, a1))
}
