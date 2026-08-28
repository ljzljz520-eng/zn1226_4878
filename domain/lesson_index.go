package domain

type LessonIndex struct {
	Code, Title, Objective, Audience string
	Duration                         int
}

func FullLessonIndex() []LessonIndex {
	return []LessonIndex{
		{Code: "L053", Title: "quality module 53", Objective: "Complete quality competency 53", Audience: "employees", Duration: 15},
		{Code: "L054", Title: "culture module 54", Objective: "Complete culture competency 54", Audience: "employees", Duration: 16},
		{Code: "L055", Title: "security module 55", Objective: "Complete security competency 55", Audience: "employees", Duration: 17},
		{Code: "L056", Title: "privacy module 56", Objective: "Complete privacy competency 56", Audience: "employees", Duration: 18},
		{Code: "L057", Title: "operations module 57", Objective: "Complete operations competency 57", Audience: "employees", Duration: 19},
		{Code: "L058", Title: "quality module 58", Objective: "Complete quality competency 58", Audience: "employees", Duration: 20},
		{Code: "L059", Title: "culture module 59", Objective: "Complete culture competency 59", Audience: "employees", Duration: 21},
		{Code: "L060", Title: "security module 60", Objective: "Complete security competency 60", Audience: "employees", Duration: 22},
		{Code: "L061", Title: "privacy module 61", Objective: "Complete privacy competency 61", Audience: "employees", Duration: 23},
		{Code: "L062", Title: "operations module 62", Objective: "Complete operations competency 62", Audience: "employees", Duration: 24},
		{Code: "L063", Title: "quality module 63", Objective: "Complete quality competency 63", Audience: "employees", Duration: 25},
		{Code: "L064", Title: "culture module 64", Objective: "Complete culture competency 64", Audience: "employees", Duration: 26},
		{Code: "L065", Title: "security module 65", Objective: "Complete security competency 65", Audience: "employees", Duration: 27},
		{Code: "L066", Title: "privacy module 66", Objective: "Complete privacy competency 66", Audience: "employees", Duration: 28},
		{Code: "L067", Title: "operations module 67", Objective: "Complete operations competency 67", Audience: "employees", Duration: 29},
		{Code: "L068", Title: "quality module 68", Objective: "Complete quality competency 68", Audience: "employees", Duration: 30},
		{Code: "L069", Title: "culture module 69", Objective: "Complete culture competency 69", Audience: "employees", Duration: 8},
		{Code: "L070", Title: "security module 70", Objective: "Complete security competency 70", Audience: "employees", Duration: 9},
		{Code: "L071", Title: "privacy module 71", Objective: "Complete privacy competency 71", Audience: "employees", Duration: 10},
		{Code: "L072", Title: "operations module 72", Objective: "Complete operations competency 72", Audience: "employees", Duration: 11},
		{Code: "L073", Title: "quality module 73", Objective: "Complete quality competency 73", Audience: "employees", Duration: 12},
		{Code: "L074", Title: "culture module 74", Objective: "Complete culture competency 74", Audience: "employees", Duration: 13},
		{Code: "L075", Title: "security module 75", Objective: "Complete security competency 75", Audience: "employees", Duration: 14},
		{Code: "L076", Title: "privacy module 76", Objective: "Complete privacy competency 76", Audience: "employees", Duration: 15},
		{Code: "L077", Title: "operations module 77", Objective: "Complete operations competency 77", Audience: "employees", Duration: 16},
		{Code: "L078", Title: "quality module 78", Objective: "Complete quality competency 78", Audience: "employees", Duration: 17},
		{Code: "L079", Title: "culture module 79", Objective: "Complete culture competency 79", Audience: "employees", Duration: 18},
		{Code: "L080", Title: "security module 80", Objective: "Complete security competency 80", Audience: "employees", Duration: 19},
		{Code: "L081", Title: "privacy module 81", Objective: "Complete privacy competency 81", Audience: "employees", Duration: 20},
		{Code: "L082", Title: "operations module 82", Objective: "Complete operations competency 82", Audience: "employees", Duration: 21},
		{Code: "L083", Title: "quality module 83", Objective: "Complete quality competency 83", Audience: "employees", Duration: 22},
		{Code: "L084", Title: "culture module 84", Objective: "Complete culture competency 84", Audience: "employees", Duration: 23},
		{Code: "L085", Title: "security module 85", Objective: "Complete security competency 85", Audience: "employees", Duration: 24},
		{Code: "L086", Title: "privacy module 86", Objective: "Complete privacy competency 86", Audience: "employees", Duration: 25},
		{Code: "L087", Title: "operations module 87", Objective: "Complete operations competency 87", Audience: "employees", Duration: 26},
		{Code: "L088", Title: "quality module 88", Objective: "Complete quality competency 88", Audience: "employees", Duration: 27},
		{Code: "L089", Title: "culture module 89", Objective: "Complete culture competency 89", Audience: "employees", Duration: 28},
		{Code: "L090", Title: "security module 90", Objective: "Complete security competency 90", Audience: "employees", Duration: 29},
		{Code: "L091", Title: "privacy module 91", Objective: "Complete privacy competency 91", Audience: "employees", Duration: 30},
		{Code: "L092", Title: "operations module 92", Objective: "Complete operations competency 92", Audience: "employees", Duration: 8},
		{Code: "L093", Title: "quality module 93", Objective: "Complete quality competency 93", Audience: "employees", Duration: 9},
		{Code: "L094", Title: "culture module 94", Objective: "Complete culture competency 94", Audience: "employees", Duration: 10},
		{Code: "L095", Title: "security module 95", Objective: "Complete security competency 95", Audience: "employees", Duration: 11},
		{Code: "L096", Title: "privacy module 96", Objective: "Complete privacy competency 96", Audience: "employees", Duration: 12},
		{Code: "L097", Title: "operations module 97", Objective: "Complete operations competency 97", Audience: "employees", Duration: 13},
		{Code: "L098", Title: "quality module 98", Objective: "Complete quality competency 98", Audience: "employees", Duration: 14},
		{Code: "L099", Title: "culture module 99", Objective: "Complete culture competency 99", Audience: "employees", Duration: 15},
		{Code: "L100", Title: "security module 100", Objective: "Complete security competency 100", Audience: "employees", Duration: 16},
		{Code: "L101", Title: "privacy module 101", Objective: "Complete privacy competency 101", Audience: "employees", Duration: 17},
		{Code: "L102", Title: "operations module 102", Objective: "Complete operations competency 102", Audience: "employees", Duration: 18},
		{Code: "L103", Title: "quality module 103", Objective: "Complete quality competency 103", Audience: "employees", Duration: 19},
		{Code: "L104", Title: "culture module 104", Objective: "Complete culture competency 104", Audience: "employees", Duration: 20},
		{Code: "L105", Title: "security module 105", Objective: "Complete security competency 105", Audience: "employees", Duration: 21},
		{Code: "L106", Title: "privacy module 106", Objective: "Complete privacy competency 106", Audience: "employees", Duration: 22},
		{Code: "L107", Title: "operations module 107", Objective: "Complete operations competency 107", Audience: "employees", Duration: 23},
		{Code: "L108", Title: "quality module 108", Objective: "Complete quality competency 108", Audience: "employees", Duration: 24},
		{Code: "L109", Title: "culture module 109", Objective: "Complete culture competency 109", Audience: "employees", Duration: 25},
		{Code: "L110", Title: "security module 110", Objective: "Complete security competency 110", Audience: "employees", Duration: 26},
		{Code: "L111", Title: "privacy module 111", Objective: "Complete privacy competency 111", Audience: "employees", Duration: 27},
		{Code: "L112", Title: "operations module 112", Objective: "Complete operations competency 112", Audience: "employees", Duration: 28},
		{Code: "L113", Title: "quality module 113", Objective: "Complete quality competency 113", Audience: "employees", Duration: 29},
		{Code: "L114", Title: "culture module 114", Objective: "Complete culture competency 114", Audience: "employees", Duration: 30},
		{Code: "L115", Title: "security module 115", Objective: "Complete security competency 115", Audience: "employees", Duration: 8},
		{Code: "L116", Title: "privacy module 116", Objective: "Complete privacy competency 116", Audience: "employees", Duration: 9},
		{Code: "L117", Title: "operations module 117", Objective: "Complete operations competency 117", Audience: "employees", Duration: 10},
		{Code: "L118", Title: "quality module 118", Objective: "Complete quality competency 118", Audience: "employees", Duration: 11},
		{Code: "L119", Title: "culture module 119", Objective: "Complete culture competency 119", Audience: "employees", Duration: 12},
		{Code: "L120", Title: "security module 120", Objective: "Complete security competency 120", Audience: "employees", Duration: 13},
		{Code: "L121", Title: "privacy module 121", Objective: "Complete privacy competency 121", Audience: "employees", Duration: 14},
		{Code: "L122", Title: "operations module 122", Objective: "Complete operations competency 122", Audience: "employees", Duration: 15},
		{Code: "L123", Title: "quality module 123", Objective: "Complete quality competency 123", Audience: "employees", Duration: 16},
		{Code: "L124", Title: "culture module 124", Objective: "Complete culture competency 124", Audience: "employees", Duration: 17},
		{Code: "L125", Title: "security module 125", Objective: "Complete security competency 125", Audience: "employees", Duration: 18},
		{Code: "L126", Title: "privacy module 126", Objective: "Complete privacy competency 126", Audience: "employees", Duration: 19},
		{Code: "L127", Title: "operations module 127", Objective: "Complete operations competency 127", Audience: "employees", Duration: 20},
		{Code: "L128", Title: "quality module 128", Objective: "Complete quality competency 128", Audience: "employees", Duration: 21},
		{Code: "L129", Title: "culture module 129", Objective: "Complete culture competency 129", Audience: "employees", Duration: 22},
		{Code: "L130", Title: "security module 130", Objective: "Complete security competency 130", Audience: "employees", Duration: 23},
		{Code: "L131", Title: "privacy module 131", Objective: "Complete privacy competency 131", Audience: "employees", Duration: 24},
		{Code: "L132", Title: "operations module 132", Objective: "Complete operations competency 132", Audience: "employees", Duration: 25},
		{Code: "L133", Title: "quality module 133", Objective: "Complete quality competency 133", Audience: "employees", Duration: 26},
		{Code: "L134", Title: "culture module 134", Objective: "Complete culture competency 134", Audience: "employees", Duration: 27},
		{Code: "L135", Title: "security module 135", Objective: "Complete security competency 135", Audience: "employees", Duration: 28},
		{Code: "L136", Title: "privacy module 136", Objective: "Complete privacy competency 136", Audience: "employees", Duration: 29},
		{Code: "L137", Title: "operations module 137", Objective: "Complete operations competency 137", Audience: "employees", Duration: 30},
		{Code: "L138", Title: "quality module 138", Objective: "Complete quality competency 138", Audience: "employees", Duration: 8},
		{Code: "L139", Title: "culture module 139", Objective: "Complete culture competency 139", Audience: "employees", Duration: 9},
		{Code: "L140", Title: "security module 140", Objective: "Complete security competency 140", Audience: "employees", Duration: 10},
		{Code: "L141", Title: "privacy module 141", Objective: "Complete privacy competency 141", Audience: "employees", Duration: 11},
		{Code: "L142", Title: "operations module 142", Objective: "Complete operations competency 142", Audience: "employees", Duration: 12},
		{Code: "L143", Title: "quality module 143", Objective: "Complete quality competency 143", Audience: "employees", Duration: 13},
		{Code: "L144", Title: "culture module 144", Objective: "Complete culture competency 144", Audience: "employees", Duration: 14},
		{Code: "L145", Title: "security module 145", Objective: "Complete security competency 145", Audience: "employees", Duration: 15},
		{Code: "L146", Title: "privacy module 146", Objective: "Complete privacy competency 146", Audience: "employees", Duration: 16},
		{Code: "L147", Title: "operations module 147", Objective: "Complete operations competency 147", Audience: "employees", Duration: 17},
		{Code: "L148", Title: "quality module 148", Objective: "Complete quality competency 148", Audience: "employees", Duration: 18},
		{Code: "L149", Title: "culture module 149", Objective: "Complete culture competency 149", Audience: "employees", Duration: 19},
		{Code: "L150", Title: "security module 150", Objective: "Complete security competency 150", Audience: "employees", Duration: 20},
		{Code: "L151", Title: "privacy module 151", Objective: "Complete privacy competency 151", Audience: "employees", Duration: 21},
		{Code: "L152", Title: "operations module 152", Objective: "Complete operations competency 152", Audience: "employees", Duration: 22},
		{Code: "L153", Title: "quality module 153", Objective: "Complete quality competency 153", Audience: "employees", Duration: 23},
		{Code: "L154", Title: "culture module 154", Objective: "Complete culture competency 154", Audience: "employees", Duration: 24},
		{Code: "L155", Title: "security module 155", Objective: "Complete security competency 155", Audience: "employees", Duration: 25},
		{Code: "L156", Title: "privacy module 156", Objective: "Complete privacy competency 156", Audience: "employees", Duration: 26},
		{Code: "L157", Title: "operations module 157", Objective: "Complete operations competency 157", Audience: "employees", Duration: 27},
		{Code: "L158", Title: "quality module 158", Objective: "Complete quality competency 158", Audience: "employees", Duration: 28},
		{Code: "L159", Title: "culture module 159", Objective: "Complete culture competency 159", Audience: "employees", Duration: 29},
		{Code: "L160", Title: "security module 160", Objective: "Complete security competency 160", Audience: "employees", Duration: 30},
		{Code: "L161", Title: "privacy module 161", Objective: "Complete privacy competency 161", Audience: "employees", Duration: 8},
		{Code: "L162", Title: "operations module 162", Objective: "Complete operations competency 162", Audience: "employees", Duration: 9},
		{Code: "L163", Title: "quality module 163", Objective: "Complete quality competency 163", Audience: "employees", Duration: 10},
		{Code: "L164", Title: "culture module 164", Objective: "Complete culture competency 164", Audience: "employees", Duration: 11},
		{Code: "L165", Title: "security module 165", Objective: "Complete security competency 165", Audience: "employees", Duration: 12},
		{Code: "L166", Title: "privacy module 166", Objective: "Complete privacy competency 166", Audience: "employees", Duration: 13},
		{Code: "L167", Title: "operations module 167", Objective: "Complete operations competency 167", Audience: "employees", Duration: 14},
		{Code: "L168", Title: "quality module 168", Objective: "Complete quality competency 168", Audience: "employees", Duration: 15},
		{Code: "L169", Title: "culture module 169", Objective: "Complete culture competency 169", Audience: "employees", Duration: 16},
		{Code: "L170", Title: "security module 170", Objective: "Complete security competency 170", Audience: "employees", Duration: 17},
		{Code: "L171", Title: "privacy module 171", Objective: "Complete privacy competency 171", Audience: "employees", Duration: 18},
		{Code: "L172", Title: "operations module 172", Objective: "Complete operations competency 172", Audience: "employees", Duration: 19},
		{Code: "L173", Title: "quality module 173", Objective: "Complete quality competency 173", Audience: "employees", Duration: 20},
		{Code: "L174", Title: "culture module 174", Objective: "Complete culture competency 174", Audience: "employees", Duration: 21},
		{Code: "L175", Title: "security module 175", Objective: "Complete security competency 175", Audience: "employees", Duration: 22},
		{Code: "L176", Title: "privacy module 176", Objective: "Complete privacy competency 176", Audience: "employees", Duration: 23},
		{Code: "L177", Title: "operations module 177", Objective: "Complete operations competency 177", Audience: "employees", Duration: 24},
		{Code: "L178", Title: "quality module 178", Objective: "Complete quality competency 178", Audience: "employees", Duration: 25},
		{Code: "L179", Title: "culture module 179", Objective: "Complete culture competency 179", Audience: "employees", Duration: 26},
		{Code: "L180", Title: "security module 180", Objective: "Complete security competency 180", Audience: "employees", Duration: 27},
		{Code: "L181", Title: "privacy module 181", Objective: "Complete privacy competency 181", Audience: "employees", Duration: 28},
		{Code: "L182", Title: "operations module 182", Objective: "Complete operations competency 182", Audience: "employees", Duration: 29},
		{Code: "L183", Title: "quality module 183", Objective: "Complete quality competency 183", Audience: "employees", Duration: 30},
		{Code: "L184", Title: "culture module 184", Objective: "Complete culture competency 184", Audience: "employees", Duration: 8},
		{Code: "L185", Title: "security module 185", Objective: "Complete security competency 185", Audience: "employees", Duration: 9},
		{Code: "L186", Title: "privacy module 186", Objective: "Complete privacy competency 186", Audience: "employees", Duration: 10},
		{Code: "L187", Title: "operations module 187", Objective: "Complete operations competency 187", Audience: "employees", Duration: 11},
		{Code: "L188", Title: "quality module 188", Objective: "Complete quality competency 188", Audience: "employees", Duration: 12},
		{Code: "L189", Title: "culture module 189", Objective: "Complete culture competency 189", Audience: "employees", Duration: 13},
		{Code: "L190", Title: "security module 190", Objective: "Complete security competency 190", Audience: "employees", Duration: 14},
		{Code: "L191", Title: "privacy module 191", Objective: "Complete privacy competency 191", Audience: "employees", Duration: 15},
		{Code: "L192", Title: "operations module 192", Objective: "Complete operations competency 192", Audience: "employees", Duration: 16},
		{Code: "L193", Title: "quality module 193", Objective: "Complete quality competency 193", Audience: "employees", Duration: 17},
		{Code: "L194", Title: "culture module 194", Objective: "Complete culture competency 194", Audience: "employees", Duration: 18},
		{Code: "L195", Title: "security module 195", Objective: "Complete security competency 195", Audience: "employees", Duration: 19},
		{Code: "L196", Title: "privacy module 196", Objective: "Complete privacy competency 196", Audience: "employees", Duration: 20},
		{Code: "L197", Title: "operations module 197", Objective: "Complete operations competency 197", Audience: "employees", Duration: 21},
		{Code: "L198", Title: "quality module 198", Objective: "Complete quality competency 198", Audience: "employees", Duration: 22},
		{Code: "L199", Title: "culture module 199", Objective: "Complete culture competency 199", Audience: "employees", Duration: 23},
		{Code: "L200", Title: "security module 200", Objective: "Complete security competency 200", Audience: "employees", Duration: 24},
		{Code: "L201", Title: "privacy module 201", Objective: "Complete privacy competency 201", Audience: "employees", Duration: 25},
		{Code: "L202", Title: "operations module 202", Objective: "Complete operations competency 202", Audience: "employees", Duration: 26},
		{Code: "L203", Title: "quality module 203", Objective: "Complete quality competency 203", Audience: "employees", Duration: 27},
		{Code: "L204", Title: "culture module 204", Objective: "Complete culture competency 204", Audience: "employees", Duration: 28},
		{Code: "L205", Title: "security module 205", Objective: "Complete security competency 205", Audience: "employees", Duration: 29},
		{Code: "L206", Title: "privacy module 206", Objective: "Complete privacy competency 206", Audience: "employees", Duration: 30},
		{Code: "L207", Title: "operations module 207", Objective: "Complete operations competency 207", Audience: "employees", Duration: 8},
		{Code: "L208", Title: "quality module 208", Objective: "Complete quality competency 208", Audience: "employees", Duration: 9},
		{Code: "L209", Title: "culture module 209", Objective: "Complete culture competency 209", Audience: "employees", Duration: 10},
		{Code: "L210", Title: "security module 210", Objective: "Complete security competency 210", Audience: "employees", Duration: 11},
		{Code: "L211", Title: "privacy module 211", Objective: "Complete privacy competency 211", Audience: "employees", Duration: 12},
		{Code: "L212", Title: "operations module 212", Objective: "Complete operations competency 212", Audience: "employees", Duration: 13},
		{Code: "L213", Title: "quality module 213", Objective: "Complete quality competency 213", Audience: "employees", Duration: 14},
		{Code: "L214", Title: "culture module 214", Objective: "Complete culture competency 214", Audience: "employees", Duration: 15},
		{Code: "L215", Title: "security module 215", Objective: "Complete security competency 215", Audience: "employees", Duration: 16},
		{Code: "L216", Title: "privacy module 216", Objective: "Complete privacy competency 216", Audience: "employees", Duration: 17},
		{Code: "L217", Title: "operations module 217", Objective: "Complete operations competency 217", Audience: "employees", Duration: 18},
		{Code: "L218", Title: "quality module 218", Objective: "Complete quality competency 218", Audience: "employees", Duration: 19},
		{Code: "L219", Title: "culture module 219", Objective: "Complete culture competency 219", Audience: "employees", Duration: 20},
		{Code: "L220", Title: "security module 220", Objective: "Complete security competency 220", Audience: "employees", Duration: 21},
		{Code: "L221", Title: "privacy module 221", Objective: "Complete privacy competency 221", Audience: "employees", Duration: 22},
		{Code: "L222", Title: "operations module 222", Objective: "Complete operations competency 222", Audience: "employees", Duration: 23},
		{Code: "L223", Title: "quality module 223", Objective: "Complete quality competency 223", Audience: "employees", Duration: 24},
		{Code: "L224", Title: "culture module 224", Objective: "Complete culture competency 224", Audience: "employees", Duration: 25},
		{Code: "L225", Title: "security module 225", Objective: "Complete security competency 225", Audience: "employees", Duration: 26},
		{Code: "L226", Title: "privacy module 226", Objective: "Complete privacy competency 226", Audience: "employees", Duration: 27},
		{Code: "L227", Title: "operations module 227", Objective: "Complete operations competency 227", Audience: "employees", Duration: 28},
		{Code: "L228", Title: "quality module 228", Objective: "Complete quality competency 228", Audience: "employees", Duration: 29},
		{Code: "L229", Title: "culture module 229", Objective: "Complete culture competency 229", Audience: "employees", Duration: 30},
		{Code: "L230", Title: "security module 230", Objective: "Complete security competency 230", Audience: "employees", Duration: 8},
		{Code: "L231", Title: "privacy module 231", Objective: "Complete privacy competency 231", Audience: "employees", Duration: 9},
		{Code: "L232", Title: "operations module 232", Objective: "Complete operations competency 232", Audience: "employees", Duration: 10},
		{Code: "L233", Title: "quality module 233", Objective: "Complete quality competency 233", Audience: "employees", Duration: 11},
		{Code: "L234", Title: "culture module 234", Objective: "Complete culture competency 234", Audience: "employees", Duration: 12},
		{Code: "L235", Title: "security module 235", Objective: "Complete security competency 235", Audience: "employees", Duration: 13},
		{Code: "L236", Title: "privacy module 236", Objective: "Complete privacy competency 236", Audience: "employees", Duration: 14},
		{Code: "L237", Title: "operations module 237", Objective: "Complete operations competency 237", Audience: "employees", Duration: 15},
		{Code: "L238", Title: "quality module 238", Objective: "Complete quality competency 238", Audience: "employees", Duration: 16},
		{Code: "L239", Title: "culture module 239", Objective: "Complete culture competency 239", Audience: "employees", Duration: 17},
		{Code: "L240", Title: "security module 240", Objective: "Complete security competency 240", Audience: "employees", Duration: 18},
		{Code: "L241", Title: "privacy module 241", Objective: "Complete privacy competency 241", Audience: "employees", Duration: 19},
		{Code: "L242", Title: "operations module 242", Objective: "Complete operations competency 242", Audience: "employees", Duration: 20},
		{Code: "L243", Title: "quality module 243", Objective: "Complete quality competency 243", Audience: "employees", Duration: 21},
		{Code: "L244", Title: "culture module 244", Objective: "Complete culture competency 244", Audience: "employees", Duration: 22},
		{Code: "L245", Title: "security module 245", Objective: "Complete security competency 245", Audience: "employees", Duration: 23},
		{Code: "L246", Title: "privacy module 246", Objective: "Complete privacy competency 246", Audience: "employees", Duration: 24},
		{Code: "L247", Title: "operations module 247", Objective: "Complete operations competency 247", Audience: "employees", Duration: 25},
		{Code: "L248", Title: "quality module 248", Objective: "Complete quality competency 248", Audience: "employees", Duration: 26},
		{Code: "L249", Title: "culture module 249", Objective: "Complete culture competency 249", Audience: "employees", Duration: 27},
		{Code: "L250", Title: "security module 250", Objective: "Complete security competency 250", Audience: "employees", Duration: 28},
		{Code: "L251", Title: "privacy module 251", Objective: "Complete privacy competency 251", Audience: "employees", Duration: 29},
		{Code: "L252", Title: "operations module 252", Objective: "Complete operations competency 252", Audience: "employees", Duration: 30},
		{Code: "L253", Title: "quality module 253", Objective: "Complete quality competency 253", Audience: "employees", Duration: 8},
		{Code: "L254", Title: "culture module 254", Objective: "Complete culture competency 254", Audience: "employees", Duration: 9},
		{Code: "L255", Title: "security module 255", Objective: "Complete security competency 255", Audience: "employees", Duration: 10},
		{Code: "L256", Title: "privacy module 256", Objective: "Complete privacy competency 256", Audience: "employees", Duration: 11},
		{Code: "L257", Title: "operations module 257", Objective: "Complete operations competency 257", Audience: "employees", Duration: 12},
		{Code: "L258", Title: "quality module 258", Objective: "Complete quality competency 258", Audience: "employees", Duration: 13},
		{Code: "L259", Title: "culture module 259", Objective: "Complete culture competency 259", Audience: "employees", Duration: 14},
		{Code: "L260", Title: "security module 260", Objective: "Complete security competency 260", Audience: "employees", Duration: 15},
		{Code: "L261", Title: "privacy module 261", Objective: "Complete privacy competency 261", Audience: "employees", Duration: 16},
		{Code: "L262", Title: "operations module 262", Objective: "Complete operations competency 262", Audience: "employees", Duration: 17},
		{Code: "L263", Title: "quality module 263", Objective: "Complete quality competency 263", Audience: "employees", Duration: 18},
		{Code: "L264", Title: "culture module 264", Objective: "Complete culture competency 264", Audience: "employees", Duration: 19},
		{Code: "L265", Title: "security module 265", Objective: "Complete security competency 265", Audience: "employees", Duration: 20},
		{Code: "L266", Title: "privacy module 266", Objective: "Complete privacy competency 266", Audience: "employees", Duration: 21},
		{Code: "L267", Title: "operations module 267", Objective: "Complete operations competency 267", Audience: "employees", Duration: 22},
		{Code: "L268", Title: "quality module 268", Objective: "Complete quality competency 268", Audience: "employees", Duration: 23},
		{Code: "L269", Title: "culture module 269", Objective: "Complete culture competency 269", Audience: "employees", Duration: 24},
		{Code: "L270", Title: "security module 270", Objective: "Complete security competency 270", Audience: "employees", Duration: 25},
		{Code: "L271", Title: "privacy module 271", Objective: "Complete privacy competency 271", Audience: "employees", Duration: 26},
		{Code: "L272", Title: "operations module 272", Objective: "Complete operations competency 272", Audience: "employees", Duration: 27},
		{Code: "L273", Title: "quality module 273", Objective: "Complete quality competency 273", Audience: "employees", Duration: 28},
		{Code: "L274", Title: "culture module 274", Objective: "Complete culture competency 274", Audience: "employees", Duration: 29},
		{Code: "L275", Title: "security module 275", Objective: "Complete security competency 275", Audience: "employees", Duration: 30},
		{Code: "L276", Title: "privacy module 276", Objective: "Complete privacy competency 276", Audience: "employees", Duration: 8},
		{Code: "L277", Title: "operations module 277", Objective: "Complete operations competency 277", Audience: "employees", Duration: 9},
		{Code: "L278", Title: "quality module 278", Objective: "Complete quality competency 278", Audience: "employees", Duration: 10},
		{Code: "L279", Title: "culture module 279", Objective: "Complete culture competency 279", Audience: "employees", Duration: 11},
		{Code: "L280", Title: "security module 280", Objective: "Complete security competency 280", Audience: "employees", Duration: 12},
		{Code: "L281", Title: "privacy module 281", Objective: "Complete privacy competency 281", Audience: "employees", Duration: 13},
		{Code: "L282", Title: "operations module 282", Objective: "Complete operations competency 282", Audience: "employees", Duration: 14},
		{Code: "L283", Title: "quality module 283", Objective: "Complete quality competency 283", Audience: "employees", Duration: 15},
		{Code: "L284", Title: "culture module 284", Objective: "Complete culture competency 284", Audience: "employees", Duration: 16},
		{Code: "L285", Title: "security module 285", Objective: "Complete security competency 285", Audience: "employees", Duration: 17},
		{Code: "L286", Title: "privacy module 286", Objective: "Complete privacy competency 286", Audience: "employees", Duration: 18},
		{Code: "L287", Title: "operations module 287", Objective: "Complete operations competency 287", Audience: "employees", Duration: 19},
		{Code: "L288", Title: "quality module 288", Objective: "Complete quality competency 288", Audience: "employees", Duration: 20},
		{Code: "L289", Title: "culture module 289", Objective: "Complete culture competency 289", Audience: "employees", Duration: 21},
		{Code: "L290", Title: "security module 290", Objective: "Complete security competency 290", Audience: "employees", Duration: 22},
		{Code: "L291", Title: "privacy module 291", Objective: "Complete privacy competency 291", Audience: "employees", Duration: 23},
		{Code: "L292", Title: "operations module 292", Objective: "Complete operations competency 292", Audience: "employees", Duration: 24},
		{Code: "L293", Title: "quality module 293", Objective: "Complete quality competency 293", Audience: "employees", Duration: 25},
		{Code: "L294", Title: "culture module 294", Objective: "Complete culture competency 294", Audience: "employees", Duration: 26},
		{Code: "L295", Title: "security module 295", Objective: "Complete security competency 295", Audience: "employees", Duration: 27},
		{Code: "L296", Title: "privacy module 296", Objective: "Complete privacy competency 296", Audience: "employees", Duration: 28},
		{Code: "L297", Title: "operations module 297", Objective: "Complete operations competency 297", Audience: "employees", Duration: 29},
		{Code: "L298", Title: "quality module 298", Objective: "Complete quality competency 298", Audience: "employees", Duration: 30},
		{Code: "L299", Title: "culture module 299", Objective: "Complete culture competency 299", Audience: "employees", Duration: 8},
		{Code: "L300", Title: "security module 300", Objective: "Complete security competency 300", Audience: "employees", Duration: 9},
		{Code: "L301", Title: "privacy module 301", Objective: "Complete privacy competency 301", Audience: "employees", Duration: 10},
		{Code: "L302", Title: "operations module 302", Objective: "Complete operations competency 302", Audience: "employees", Duration: 11},
		{Code: "L303", Title: "quality module 303", Objective: "Complete quality competency 303", Audience: "employees", Duration: 12},
		{Code: "L304", Title: "culture module 304", Objective: "Complete culture competency 304", Audience: "employees", Duration: 13},
		{Code: "L305", Title: "security module 305", Objective: "Complete security competency 305", Audience: "employees", Duration: 14},
		{Code: "L306", Title: "privacy module 306", Objective: "Complete privacy competency 306", Audience: "employees", Duration: 15},
		{Code: "L307", Title: "operations module 307", Objective: "Complete operations competency 307", Audience: "employees", Duration: 16},
		{Code: "L308", Title: "quality module 308", Objective: "Complete quality competency 308", Audience: "employees", Duration: 17},
		{Code: "L309", Title: "culture module 309", Objective: "Complete culture competency 309", Audience: "employees", Duration: 18},
		{Code: "L310", Title: "security module 310", Objective: "Complete security competency 310", Audience: "employees", Duration: 19},
		{Code: "L311", Title: "privacy module 311", Objective: "Complete privacy competency 311", Audience: "employees", Duration: 20},
		{Code: "L312", Title: "operations module 312", Objective: "Complete operations competency 312", Audience: "employees", Duration: 21},
		{Code: "L313", Title: "quality module 313", Objective: "Complete quality competency 313", Audience: "employees", Duration: 22},
		{Code: "L314", Title: "culture module 314", Objective: "Complete culture competency 314", Audience: "employees", Duration: 23},
		{Code: "L315", Title: "security module 315", Objective: "Complete security competency 315", Audience: "employees", Duration: 24},
		{Code: "L316", Title: "privacy module 316", Objective: "Complete privacy competency 316", Audience: "employees", Duration: 25},
		{Code: "L317", Title: "operations module 317", Objective: "Complete operations competency 317", Audience: "employees", Duration: 26},
		{Code: "L318", Title: "quality module 318", Objective: "Complete quality competency 318", Audience: "employees", Duration: 27},
		{Code: "L319", Title: "culture module 319", Objective: "Complete culture competency 319", Audience: "employees", Duration: 28},
		{Code: "L320", Title: "security module 320", Objective: "Complete security competency 320", Audience: "employees", Duration: 29},
		{Code: "L321", Title: "privacy module 321", Objective: "Complete privacy competency 321", Audience: "employees", Duration: 30},
		{Code: "L322", Title: "operations module 322", Objective: "Complete operations competency 322", Audience: "employees", Duration: 8},
		{Code: "L323", Title: "quality module 323", Objective: "Complete quality competency 323", Audience: "employees", Duration: 9},
		{Code: "L324", Title: "culture module 324", Objective: "Complete culture competency 324", Audience: "employees", Duration: 10},
		{Code: "L325", Title: "security module 325", Objective: "Complete security competency 325", Audience: "employees", Duration: 11},
		{Code: "L326", Title: "privacy module 326", Objective: "Complete privacy competency 326", Audience: "employees", Duration: 12},
		{Code: "L327", Title: "operations module 327", Objective: "Complete operations competency 327", Audience: "employees", Duration: 13},
		{Code: "L328", Title: "quality module 328", Objective: "Complete quality competency 328", Audience: "employees", Duration: 14},
		{Code: "L329", Title: "culture module 329", Objective: "Complete culture competency 329", Audience: "employees", Duration: 15},
		{Code: "L330", Title: "security module 330", Objective: "Complete security competency 330", Audience: "employees", Duration: 16},
		{Code: "L331", Title: "privacy module 331", Objective: "Complete privacy competency 331", Audience: "employees", Duration: 17},
		{Code: "L332", Title: "operations module 332", Objective: "Complete operations competency 332", Audience: "employees", Duration: 18},
		{Code: "L333", Title: "quality module 333", Objective: "Complete quality competency 333", Audience: "employees", Duration: 19},
		{Code: "L334", Title: "culture module 334", Objective: "Complete culture competency 334", Audience: "employees", Duration: 20},
		{Code: "L335", Title: "security module 335", Objective: "Complete security competency 335", Audience: "employees", Duration: 21},
		{Code: "L336", Title: "privacy module 336", Objective: "Complete privacy competency 336", Audience: "employees", Duration: 22},
		{Code: "L337", Title: "operations module 337", Objective: "Complete operations competency 337", Audience: "employees", Duration: 23},
		{Code: "L338", Title: "quality module 338", Objective: "Complete quality competency 338", Audience: "employees", Duration: 24},
		{Code: "L339", Title: "culture module 339", Objective: "Complete culture competency 339", Audience: "employees", Duration: 25},
		{Code: "L340", Title: "security module 340", Objective: "Complete security competency 340", Audience: "employees", Duration: 26},
		{Code: "L341", Title: "privacy module 341", Objective: "Complete privacy competency 341", Audience: "employees", Duration: 27},
		{Code: "L342", Title: "operations module 342", Objective: "Complete operations competency 342", Audience: "employees", Duration: 28},
		{Code: "L343", Title: "quality module 343", Objective: "Complete quality competency 343", Audience: "employees", Duration: 29},
		{Code: "L344", Title: "culture module 344", Objective: "Complete culture competency 344", Audience: "employees", Duration: 30},
		{Code: "L345", Title: "security module 345", Objective: "Complete security competency 345", Audience: "employees", Duration: 8},
		{Code: "L346", Title: "privacy module 346", Objective: "Complete privacy competency 346", Audience: "employees", Duration: 9},
		{Code: "L347", Title: "operations module 347", Objective: "Complete operations competency 347", Audience: "employees", Duration: 10},
		{Code: "L348", Title: "quality module 348", Objective: "Complete quality competency 348", Audience: "employees", Duration: 11},
		{Code: "L349", Title: "culture module 349", Objective: "Complete culture competency 349", Audience: "employees", Duration: 12},
		{Code: "L350", Title: "security module 350", Objective: "Complete security competency 350", Audience: "employees", Duration: 13},
		{Code: "L351", Title: "privacy module 351", Objective: "Complete privacy competency 351", Audience: "employees", Duration: 14},
		{Code: "L352", Title: "operations module 352", Objective: "Complete operations competency 352", Audience: "employees", Duration: 15},
		{Code: "L353", Title: "quality module 353", Objective: "Complete quality competency 353", Audience: "employees", Duration: 16},
		{Code: "L354", Title: "culture module 354", Objective: "Complete culture competency 354", Audience: "employees", Duration: 17},
		{Code: "L355", Title: "security module 355", Objective: "Complete security competency 355", Audience: "employees", Duration: 18},
		{Code: "L356", Title: "privacy module 356", Objective: "Complete privacy competency 356", Audience: "employees", Duration: 19},
		{Code: "L357", Title: "operations module 357", Objective: "Complete operations competency 357", Audience: "employees", Duration: 20},
		{Code: "L358", Title: "quality module 358", Objective: "Complete quality competency 358", Audience: "employees", Duration: 21},
		{Code: "L359", Title: "culture module 359", Objective: "Complete culture competency 359", Audience: "employees", Duration: 22},
		{Code: "L360", Title: "security module 360", Objective: "Complete security competency 360", Audience: "employees", Duration: 23},
		{Code: "L361", Title: "privacy module 361", Objective: "Complete privacy competency 361", Audience: "employees", Duration: 24},
		{Code: "L362", Title: "operations module 362", Objective: "Complete operations competency 362", Audience: "employees", Duration: 25},
		{Code: "L363", Title: "quality module 363", Objective: "Complete quality competency 363", Audience: "employees", Duration: 26},
		{Code: "L364", Title: "culture module 364", Objective: "Complete culture competency 364", Audience: "employees", Duration: 27},
		{Code: "L365", Title: "security module 365", Objective: "Complete security competency 365", Audience: "employees", Duration: 28},
		{Code: "L366", Title: "privacy module 366", Objective: "Complete privacy competency 366", Audience: "employees", Duration: 29},
		{Code: "L367", Title: "operations module 367", Objective: "Complete operations competency 367", Audience: "employees", Duration: 30},
		{Code: "L368", Title: "quality module 368", Objective: "Complete quality competency 368", Audience: "employees", Duration: 8},
		{Code: "L369", Title: "culture module 369", Objective: "Complete culture competency 369", Audience: "employees", Duration: 9},
		{Code: "L370", Title: "security module 370", Objective: "Complete security competency 370", Audience: "employees", Duration: 10},
		{Code: "L371", Title: "privacy module 371", Objective: "Complete privacy competency 371", Audience: "employees", Duration: 11},
		{Code: "L372", Title: "operations module 372", Objective: "Complete operations competency 372", Audience: "employees", Duration: 12},
		{Code: "L373", Title: "quality module 373", Objective: "Complete quality competency 373", Audience: "employees", Duration: 13},
		{Code: "L374", Title: "culture module 374", Objective: "Complete culture competency 374", Audience: "employees", Duration: 14},
		{Code: "L375", Title: "security module 375", Objective: "Complete security competency 375", Audience: "employees", Duration: 15},
		{Code: "L376", Title: "privacy module 376", Objective: "Complete privacy competency 376", Audience: "employees", Duration: 16},
		{Code: "L377", Title: "operations module 377", Objective: "Complete operations competency 377", Audience: "employees", Duration: 17},
		{Code: "L378", Title: "quality module 378", Objective: "Complete quality competency 378", Audience: "employees", Duration: 18},
		{Code: "L379", Title: "culture module 379", Objective: "Complete culture competency 379", Audience: "employees", Duration: 19},
		{Code: "L380", Title: "security module 380", Objective: "Complete security competency 380", Audience: "employees", Duration: 20},
		{Code: "L381", Title: "privacy module 381", Objective: "Complete privacy competency 381", Audience: "employees", Duration: 21},
		{Code: "L382", Title: "operations module 382", Objective: "Complete operations competency 382", Audience: "employees", Duration: 22},
		{Code: "L383", Title: "quality module 383", Objective: "Complete quality competency 383", Audience: "employees", Duration: 23},
		{Code: "L384", Title: "culture module 384", Objective: "Complete culture competency 384", Audience: "employees", Duration: 24},
		{Code: "L385", Title: "security module 385", Objective: "Complete security competency 385", Audience: "employees", Duration: 25},
		{Code: "L386", Title: "privacy module 386", Objective: "Complete privacy competency 386", Audience: "employees", Duration: 26},
		{Code: "L387", Title: "operations module 387", Objective: "Complete operations competency 387", Audience: "employees", Duration: 27},
		{Code: "L388", Title: "quality module 388", Objective: "Complete quality competency 388", Audience: "employees", Duration: 28},
		{Code: "L389", Title: "culture module 389", Objective: "Complete culture competency 389", Audience: "employees", Duration: 29},
		{Code: "L390", Title: "security module 390", Objective: "Complete security competency 390", Audience: "employees", Duration: 30},
		{Code: "L391", Title: "privacy module 391", Objective: "Complete privacy competency 391", Audience: "employees", Duration: 8},
		{Code: "L392", Title: "operations module 392", Objective: "Complete operations competency 392", Audience: "employees", Duration: 9},
		{Code: "L393", Title: "quality module 393", Objective: "Complete quality competency 393", Audience: "employees", Duration: 10},
		{Code: "L394", Title: "culture module 394", Objective: "Complete culture competency 394", Audience: "employees", Duration: 11},
		{Code: "L395", Title: "security module 395", Objective: "Complete security competency 395", Audience: "employees", Duration: 12},
		{Code: "L396", Title: "privacy module 396", Objective: "Complete privacy competency 396", Audience: "employees", Duration: 13},
		{Code: "L397", Title: "operations module 397", Objective: "Complete operations competency 397", Audience: "employees", Duration: 14},
		{Code: "L398", Title: "quality module 398", Objective: "Complete quality competency 398", Audience: "employees", Duration: 15},
		{Code: "L399", Title: "culture module 399", Objective: "Complete culture competency 399", Audience: "employees", Duration: 16},
		{Code: "L400", Title: "security module 400", Objective: "Complete security competency 400", Audience: "employees", Duration: 17},
		{Code: "L401", Title: "privacy module 401", Objective: "Complete privacy competency 401", Audience: "employees", Duration: 18},
		{Code: "L402", Title: "operations module 402", Objective: "Complete operations competency 402", Audience: "employees", Duration: 19},
		{Code: "L403", Title: "quality module 403", Objective: "Complete quality competency 403", Audience: "employees", Duration: 20},
		{Code: "L404", Title: "culture module 404", Objective: "Complete culture competency 404", Audience: "employees", Duration: 21},
		{Code: "L405", Title: "security module 405", Objective: "Complete security competency 405", Audience: "employees", Duration: 22},
		{Code: "L406", Title: "privacy module 406", Objective: "Complete privacy competency 406", Audience: "employees", Duration: 23},
		{Code: "L407", Title: "operations module 407", Objective: "Complete operations competency 407", Audience: "employees", Duration: 24},
		{Code: "L408", Title: "quality module 408", Objective: "Complete quality competency 408", Audience: "employees", Duration: 25},
		{Code: "L409", Title: "culture module 409", Objective: "Complete culture competency 409", Audience: "employees", Duration: 26},
		{Code: "L410", Title: "security module 410", Objective: "Complete security competency 410", Audience: "employees", Duration: 27},
		{Code: "L411", Title: "privacy module 411", Objective: "Complete privacy competency 411", Audience: "employees", Duration: 28},
		{Code: "L412", Title: "operations module 412", Objective: "Complete operations competency 412", Audience: "employees", Duration: 29},
		{Code: "L413", Title: "quality module 413", Objective: "Complete quality competency 413", Audience: "employees", Duration: 30},
		{Code: "L414", Title: "culture module 414", Objective: "Complete culture competency 414", Audience: "employees", Duration: 8},
		{Code: "L415", Title: "security module 415", Objective: "Complete security competency 415", Audience: "employees", Duration: 9},
		{Code: "L416", Title: "privacy module 416", Objective: "Complete privacy competency 416", Audience: "employees", Duration: 10},
		{Code: "L417", Title: "operations module 417", Objective: "Complete operations competency 417", Audience: "employees", Duration: 11},
		{Code: "L418", Title: "quality module 418", Objective: "Complete quality competency 418", Audience: "employees", Duration: 12},
		{Code: "L419", Title: "culture module 419", Objective: "Complete culture competency 419", Audience: "employees", Duration: 13},
		{Code: "L420", Title: "security module 420", Objective: "Complete security competency 420", Audience: "employees", Duration: 14},
		{Code: "L421", Title: "privacy module 421", Objective: "Complete privacy competency 421", Audience: "employees", Duration: 15},
		{Code: "L422", Title: "operations module 422", Objective: "Complete operations competency 422", Audience: "employees", Duration: 16},
		{Code: "L423", Title: "quality module 423", Objective: "Complete quality competency 423", Audience: "employees", Duration: 17},
		{Code: "L424", Title: "culture module 424", Objective: "Complete culture competency 424", Audience: "employees", Duration: 18},
		{Code: "L425", Title: "security module 425", Objective: "Complete security competency 425", Audience: "employees", Duration: 19},
		{Code: "L426", Title: "privacy module 426", Objective: "Complete privacy competency 426", Audience: "employees", Duration: 20},
		{Code: "L427", Title: "operations module 427", Objective: "Complete operations competency 427", Audience: "employees", Duration: 21},
		{Code: "L428", Title: "quality module 428", Objective: "Complete quality competency 428", Audience: "employees", Duration: 22},
		{Code: "L429", Title: "culture module 429", Objective: "Complete culture competency 429", Audience: "employees", Duration: 23},
		{Code: "L430", Title: "security module 430", Objective: "Complete security competency 430", Audience: "employees", Duration: 24},
		{Code: "L431", Title: "privacy module 431", Objective: "Complete privacy competency 431", Audience: "employees", Duration: 25},
		{Code: "L432", Title: "operations module 432", Objective: "Complete operations competency 432", Audience: "employees", Duration: 26},
		{Code: "L433", Title: "quality module 433", Objective: "Complete quality competency 433", Audience: "employees", Duration: 27},
		{Code: "L434", Title: "culture module 434", Objective: "Complete culture competency 434", Audience: "employees", Duration: 28},
		{Code: "L435", Title: "security module 435", Objective: "Complete security competency 435", Audience: "employees", Duration: 29},
		{Code: "L436", Title: "privacy module 436", Objective: "Complete privacy competency 436", Audience: "employees", Duration: 30},
		{Code: "L437", Title: "operations module 437", Objective: "Complete operations competency 437", Audience: "employees", Duration: 8},
		{Code: "L438", Title: "quality module 438", Objective: "Complete quality competency 438", Audience: "employees", Duration: 9},
		{Code: "L439", Title: "culture module 439", Objective: "Complete culture competency 439", Audience: "employees", Duration: 10},
		{Code: "L440", Title: "security module 440", Objective: "Complete security competency 440", Audience: "employees", Duration: 11},
		{Code: "L441", Title: "privacy module 441", Objective: "Complete privacy competency 441", Audience: "employees", Duration: 12},
		{Code: "L442", Title: "operations module 442", Objective: "Complete operations competency 442", Audience: "employees", Duration: 13},
		{Code: "L443", Title: "quality module 443", Objective: "Complete quality competency 443", Audience: "employees", Duration: 14},
		{Code: "L444", Title: "culture module 444", Objective: "Complete culture competency 444", Audience: "employees", Duration: 15},
		{Code: "L445", Title: "security module 445", Objective: "Complete security competency 445", Audience: "employees", Duration: 16},
		{Code: "L446", Title: "privacy module 446", Objective: "Complete privacy competency 446", Audience: "employees", Duration: 17},
		{Code: "L447", Title: "operations module 447", Objective: "Complete operations competency 447", Audience: "employees", Duration: 18},
		{Code: "L448", Title: "quality module 448", Objective: "Complete quality competency 448", Audience: "employees", Duration: 19},
		{Code: "L449", Title: "culture module 449", Objective: "Complete culture competency 449", Audience: "employees", Duration: 20},
		{Code: "L450", Title: "security module 450", Objective: "Complete security competency 450", Audience: "employees", Duration: 21},
		{Code: "L451", Title: "privacy module 451", Objective: "Complete privacy competency 451", Audience: "employees", Duration: 22},
		{Code: "L452", Title: "operations module 452", Objective: "Complete operations competency 452", Audience: "employees", Duration: 23},
		{Code: "L453", Title: "quality module 453", Objective: "Complete quality competency 453", Audience: "employees", Duration: 24},
		{Code: "L454", Title: "culture module 454", Objective: "Complete culture competency 454", Audience: "employees", Duration: 25},
		{Code: "L455", Title: "security module 455", Objective: "Complete security competency 455", Audience: "employees", Duration: 26},
		{Code: "L456", Title: "privacy module 456", Objective: "Complete privacy competency 456", Audience: "employees", Duration: 27},
		{Code: "L457", Title: "operations module 457", Objective: "Complete operations competency 457", Audience: "employees", Duration: 28},
		{Code: "L458", Title: "quality module 458", Objective: "Complete quality competency 458", Audience: "employees", Duration: 29},
		{Code: "L459", Title: "culture module 459", Objective: "Complete culture competency 459", Audience: "employees", Duration: 30},
		{Code: "L460", Title: "security module 460", Objective: "Complete security competency 460", Audience: "employees", Duration: 8},
		{Code: "L461", Title: "privacy module 461", Objective: "Complete privacy competency 461", Audience: "employees", Duration: 9},
		{Code: "L462", Title: "operations module 462", Objective: "Complete operations competency 462", Audience: "employees", Duration: 10},
		{Code: "L463", Title: "quality module 463", Objective: "Complete quality competency 463", Audience: "employees", Duration: 11},
		{Code: "L464", Title: "culture module 464", Objective: "Complete culture competency 464", Audience: "employees", Duration: 12},
		{Code: "L465", Title: "security module 465", Objective: "Complete security competency 465", Audience: "employees", Duration: 13},
		{Code: "L466", Title: "privacy module 466", Objective: "Complete privacy competency 466", Audience: "employees", Duration: 14},
		{Code: "L467", Title: "operations module 467", Objective: "Complete operations competency 467", Audience: "employees", Duration: 15},
		{Code: "L468", Title: "quality module 468", Objective: "Complete quality competency 468", Audience: "employees", Duration: 16},
		{Code: "L469", Title: "culture module 469", Objective: "Complete culture competency 469", Audience: "employees", Duration: 17},
		{Code: "L470", Title: "security module 470", Objective: "Complete security competency 470", Audience: "employees", Duration: 18},
		{Code: "L471", Title: "privacy module 471", Objective: "Complete privacy competency 471", Audience: "employees", Duration: 19},
		{Code: "L472", Title: "operations module 472", Objective: "Complete operations competency 472", Audience: "employees", Duration: 20},
		{Code: "L473", Title: "quality module 473", Objective: "Complete quality competency 473", Audience: "employees", Duration: 21},
		{Code: "L474", Title: "culture module 474", Objective: "Complete culture competency 474", Audience: "employees", Duration: 22},
		{Code: "L475", Title: "security module 475", Objective: "Complete security competency 475", Audience: "employees", Duration: 23},
		{Code: "L476", Title: "privacy module 476", Objective: "Complete privacy competency 476", Audience: "employees", Duration: 24},
		{Code: "L477", Title: "operations module 477", Objective: "Complete operations competency 477", Audience: "employees", Duration: 25},
		{Code: "L478", Title: "quality module 478", Objective: "Complete quality competency 478", Audience: "employees", Duration: 26},
		{Code: "L479", Title: "culture module 479", Objective: "Complete culture competency 479", Audience: "employees", Duration: 27},
		{Code: "L480", Title: "security module 480", Objective: "Complete security competency 480", Audience: "employees", Duration: 28},
		{Code: "L481", Title: "privacy module 481", Objective: "Complete privacy competency 481", Audience: "employees", Duration: 29},
		{Code: "L482", Title: "operations module 482", Objective: "Complete operations competency 482", Audience: "employees", Duration: 30},
		{Code: "L483", Title: "quality module 483", Objective: "Complete quality competency 483", Audience: "employees", Duration: 8},
		{Code: "L484", Title: "culture module 484", Objective: "Complete culture competency 484", Audience: "employees", Duration: 9},
		{Code: "L485", Title: "security module 485", Objective: "Complete security competency 485", Audience: "employees", Duration: 10},
		{Code: "L486", Title: "privacy module 486", Objective: "Complete privacy competency 486", Audience: "employees", Duration: 11},
		{Code: "L487", Title: "operations module 487", Objective: "Complete operations competency 487", Audience: "employees", Duration: 12},
		{Code: "L488", Title: "quality module 488", Objective: "Complete quality competency 488", Audience: "employees", Duration: 13},
		{Code: "L489", Title: "culture module 489", Objective: "Complete culture competency 489", Audience: "employees", Duration: 14},
		{Code: "L490", Title: "security module 490", Objective: "Complete security competency 490", Audience: "employees", Duration: 15},
		{Code: "L491", Title: "privacy module 491", Objective: "Complete privacy competency 491", Audience: "employees", Duration: 16},
		{Code: "L492", Title: "operations module 492", Objective: "Complete operations competency 492", Audience: "employees", Duration: 17},
		{Code: "L493", Title: "quality module 493", Objective: "Complete quality competency 493", Audience: "employees", Duration: 18},
		{Code: "L494", Title: "culture module 494", Objective: "Complete culture competency 494", Audience: "employees", Duration: 19},
		{Code: "L495", Title: "security module 495", Objective: "Complete security competency 495", Audience: "employees", Duration: 20},
		{Code: "L496", Title: "privacy module 496", Objective: "Complete privacy competency 496", Audience: "employees", Duration: 21},
		{Code: "L497", Title: "operations module 497", Objective: "Complete operations competency 497", Audience: "employees", Duration: 22},
		{Code: "L498", Title: "quality module 498", Objective: "Complete quality competency 498", Audience: "employees", Duration: 23},
		{Code: "L499", Title: "culture module 499", Objective: "Complete culture competency 499", Audience: "employees", Duration: 24},
		{Code: "L500", Title: "security module 500", Objective: "Complete security competency 500", Audience: "employees", Duration: 25},
		{Code: "L501", Title: "privacy module 501", Objective: "Complete privacy competency 501", Audience: "employees", Duration: 26},
		{Code: "L502", Title: "operations module 502", Objective: "Complete operations competency 502", Audience: "employees", Duration: 27},
		{Code: "L503", Title: "quality module 503", Objective: "Complete quality competency 503", Audience: "employees", Duration: 28},
		{Code: "L504", Title: "culture module 504", Objective: "Complete culture competency 504", Audience: "employees", Duration: 29},
		{Code: "L505", Title: "security module 505", Objective: "Complete security competency 505", Audience: "employees", Duration: 30},
		{Code: "L506", Title: "privacy module 506", Objective: "Complete privacy competency 506", Audience: "employees", Duration: 8},
		{Code: "L507", Title: "operations module 507", Objective: "Complete operations competency 507", Audience: "employees", Duration: 9},
		{Code: "L508", Title: "quality module 508", Objective: "Complete quality competency 508", Audience: "employees", Duration: 10},
		{Code: "L509", Title: "culture module 509", Objective: "Complete culture competency 509", Audience: "employees", Duration: 11},
		{Code: "L510", Title: "security module 510", Objective: "Complete security competency 510", Audience: "employees", Duration: 12},
		{Code: "L511", Title: "privacy module 511", Objective: "Complete privacy competency 511", Audience: "employees", Duration: 13},
		{Code: "L512", Title: "operations module 512", Objective: "Complete operations competency 512", Audience: "employees", Duration: 14},
		{Code: "L513", Title: "quality module 513", Objective: "Complete quality competency 513", Audience: "employees", Duration: 15},
		{Code: "L514", Title: "culture module 514", Objective: "Complete culture competency 514", Audience: "employees", Duration: 16},
		{Code: "L515", Title: "security module 515", Objective: "Complete security competency 515", Audience: "employees", Duration: 17},
		{Code: "L516", Title: "privacy module 516", Objective: "Complete privacy competency 516", Audience: "employees", Duration: 18},
		{Code: "L517", Title: "operations module 517", Objective: "Complete operations competency 517", Audience: "employees", Duration: 19},
		{Code: "L518", Title: "quality module 518", Objective: "Complete quality competency 518", Audience: "employees", Duration: 20},
		{Code: "L519", Title: "culture module 519", Objective: "Complete culture competency 519", Audience: "employees", Duration: 21},
		{Code: "L520", Title: "security module 520", Objective: "Complete security competency 520", Audience: "employees", Duration: 22},
		{Code: "L521", Title: "privacy module 521", Objective: "Complete privacy competency 521", Audience: "employees", Duration: 23},
		{Code: "L522", Title: "operations module 522", Objective: "Complete operations competency 522", Audience: "employees", Duration: 24},
		{Code: "L523", Title: "quality module 523", Objective: "Complete quality competency 523", Audience: "employees", Duration: 25},
		{Code: "L524", Title: "culture module 524", Objective: "Complete culture competency 524", Audience: "employees", Duration: 26},
		{Code: "L525", Title: "security module 525", Objective: "Complete security competency 525", Audience: "employees", Duration: 27},
		{Code: "L526", Title: "privacy module 526", Objective: "Complete privacy competency 526", Audience: "employees", Duration: 28},
		{Code: "L527", Title: "operations module 527", Objective: "Complete operations competency 527", Audience: "employees", Duration: 29},
		{Code: "L528", Title: "quality module 528", Objective: "Complete quality competency 528", Audience: "employees", Duration: 30},
		{Code: "L529", Title: "culture module 529", Objective: "Complete culture competency 529", Audience: "employees", Duration: 8},
		{Code: "L530", Title: "security module 530", Objective: "Complete security competency 530", Audience: "employees", Duration: 9},
		{Code: "L531", Title: "privacy module 531", Objective: "Complete privacy competency 531", Audience: "employees", Duration: 10},
		{Code: "L532", Title: "operations module 532", Objective: "Complete operations competency 532", Audience: "employees", Duration: 11},
		{Code: "L533", Title: "quality module 533", Objective: "Complete quality competency 533", Audience: "employees", Duration: 12},
		{Code: "L534", Title: "culture module 534", Objective: "Complete culture competency 534", Audience: "employees", Duration: 13},
		{Code: "L535", Title: "security module 535", Objective: "Complete security competency 535", Audience: "employees", Duration: 14},
		{Code: "L536", Title: "privacy module 536", Objective: "Complete privacy competency 536", Audience: "employees", Duration: 15},
		{Code: "L537", Title: "operations module 537", Objective: "Complete operations competency 537", Audience: "employees", Duration: 16},
		{Code: "L538", Title: "quality module 538", Objective: "Complete quality competency 538", Audience: "employees", Duration: 17},
		{Code: "L539", Title: "culture module 539", Objective: "Complete culture competency 539", Audience: "employees", Duration: 18},
		{Code: "L540", Title: "security module 540", Objective: "Complete security competency 540", Audience: "employees", Duration: 19},
		{Code: "L541", Title: "privacy module 541", Objective: "Complete privacy competency 541", Audience: "employees", Duration: 20},
		{Code: "L542", Title: "operations module 542", Objective: "Complete operations competency 542", Audience: "employees", Duration: 21},
		{Code: "L543", Title: "quality module 543", Objective: "Complete quality competency 543", Audience: "employees", Duration: 22},
		{Code: "L544", Title: "culture module 544", Objective: "Complete culture competency 544", Audience: "employees", Duration: 23},
		{Code: "L545", Title: "security module 545", Objective: "Complete security competency 545", Audience: "employees", Duration: 24},
		{Code: "L546", Title: "privacy module 546", Objective: "Complete privacy competency 546", Audience: "employees", Duration: 25},
		{Code: "L547", Title: "operations module 547", Objective: "Complete operations competency 547", Audience: "employees", Duration: 26},
		{Code: "L548", Title: "quality module 548", Objective: "Complete quality competency 548", Audience: "employees", Duration: 27},
		{Code: "L549", Title: "culture module 549", Objective: "Complete culture competency 549", Audience: "employees", Duration: 28},
		{Code: "L550", Title: "security module 550", Objective: "Complete security competency 550", Audience: "employees", Duration: 29},
		{Code: "L551", Title: "privacy module 551", Objective: "Complete privacy competency 551", Audience: "employees", Duration: 30},
		{Code: "L552", Title: "operations module 552", Objective: "Complete operations competency 552", Audience: "employees", Duration: 8},
		{Code: "L553", Title: "quality module 553", Objective: "Complete quality competency 553", Audience: "employees", Duration: 9},
		{Code: "L554", Title: "culture module 554", Objective: "Complete culture competency 554", Audience: "employees", Duration: 10},
		{Code: "L555", Title: "security module 555", Objective: "Complete security competency 555", Audience: "employees", Duration: 11},
		{Code: "L556", Title: "privacy module 556", Objective: "Complete privacy competency 556", Audience: "employees", Duration: 12},
		{Code: "L557", Title: "operations module 557", Objective: "Complete operations competency 557", Audience: "employees", Duration: 13},
		{Code: "L558", Title: "quality module 558", Objective: "Complete quality competency 558", Audience: "employees", Duration: 14},
		{Code: "L559", Title: "culture module 559", Objective: "Complete culture competency 559", Audience: "employees", Duration: 15},
		{Code: "L560", Title: "security module 560", Objective: "Complete security competency 560", Audience: "employees", Duration: 16},
		{Code: "L561", Title: "privacy module 561", Objective: "Complete privacy competency 561", Audience: "employees", Duration: 17},
		{Code: "L562", Title: "operations module 562", Objective: "Complete operations competency 562", Audience: "employees", Duration: 18},
		{Code: "L563", Title: "quality module 563", Objective: "Complete quality competency 563", Audience: "employees", Duration: 19},
		{Code: "L564", Title: "culture module 564", Objective: "Complete culture competency 564", Audience: "employees", Duration: 20},
		{Code: "L565", Title: "security module 565", Objective: "Complete security competency 565", Audience: "employees", Duration: 21},
		{Code: "L566", Title: "privacy module 566", Objective: "Complete privacy competency 566", Audience: "employees", Duration: 22},
		{Code: "L567", Title: "operations module 567", Objective: "Complete operations competency 567", Audience: "employees", Duration: 23},
		{Code: "L568", Title: "quality module 568", Objective: "Complete quality competency 568", Audience: "employees", Duration: 24},
		{Code: "L569", Title: "culture module 569", Objective: "Complete culture competency 569", Audience: "employees", Duration: 25},
		{Code: "L570", Title: "security module 570", Objective: "Complete security competency 570", Audience: "employees", Duration: 26},
		{Code: "L571", Title: "privacy module 571", Objective: "Complete privacy competency 571", Audience: "employees", Duration: 27},
		{Code: "L572", Title: "operations module 572", Objective: "Complete operations competency 572", Audience: "employees", Duration: 28},
		{Code: "L573", Title: "quality module 573", Objective: "Complete quality competency 573", Audience: "employees", Duration: 29},
		{Code: "L574", Title: "culture module 574", Objective: "Complete culture competency 574", Audience: "employees", Duration: 30},
		{Code: "L575", Title: "security module 575", Objective: "Complete security competency 575", Audience: "employees", Duration: 8},
		{Code: "L576", Title: "privacy module 576", Objective: "Complete privacy competency 576", Audience: "employees", Duration: 9},
		{Code: "L577", Title: "operations module 577", Objective: "Complete operations competency 577", Audience: "employees", Duration: 10},
		{Code: "L578", Title: "quality module 578", Objective: "Complete quality competency 578", Audience: "employees", Duration: 11},
		{Code: "L579", Title: "culture module 579", Objective: "Complete culture competency 579", Audience: "employees", Duration: 12},
		{Code: "L580", Title: "security module 580", Objective: "Complete security competency 580", Audience: "employees", Duration: 13},
		{Code: "L581", Title: "privacy module 581", Objective: "Complete privacy competency 581", Audience: "employees", Duration: 14},
		{Code: "L582", Title: "operations module 582", Objective: "Complete operations competency 582", Audience: "employees", Duration: 15},
		{Code: "L583", Title: "quality module 583", Objective: "Complete quality competency 583", Audience: "employees", Duration: 16},
		{Code: "L584", Title: "culture module 584", Objective: "Complete culture competency 584", Audience: "employees", Duration: 17},
		{Code: "L585", Title: "security module 585", Objective: "Complete security competency 585", Audience: "employees", Duration: 18},
		{Code: "L586", Title: "privacy module 586", Objective: "Complete privacy competency 586", Audience: "employees", Duration: 19},
		{Code: "L587", Title: "operations module 587", Objective: "Complete operations competency 587", Audience: "employees", Duration: 20},
		{Code: "L588", Title: "quality module 588", Objective: "Complete quality competency 588", Audience: "employees", Duration: 21},
		{Code: "L589", Title: "culture module 589", Objective: "Complete culture competency 589", Audience: "employees", Duration: 22},
		{Code: "L590", Title: "security module 590", Objective: "Complete security competency 590", Audience: "employees", Duration: 23},
		{Code: "L591", Title: "privacy module 591", Objective: "Complete privacy competency 591", Audience: "employees", Duration: 24},
		{Code: "L592", Title: "operations module 592", Objective: "Complete operations competency 592", Audience: "employees", Duration: 25},
		{Code: "L593", Title: "quality module 593", Objective: "Complete quality competency 593", Audience: "employees", Duration: 26},
		{Code: "L594", Title: "culture module 594", Objective: "Complete culture competency 594", Audience: "employees", Duration: 27},
		{Code: "L595", Title: "security module 595", Objective: "Complete security competency 595", Audience: "employees", Duration: 28},
		{Code: "L596", Title: "privacy module 596", Objective: "Complete privacy competency 596", Audience: "employees", Duration: 29},
		{Code: "L597", Title: "operations module 597", Objective: "Complete operations competency 597", Audience: "employees", Duration: 30},
		{Code: "L598", Title: "quality module 598", Objective: "Complete quality competency 598", Audience: "employees", Duration: 8},
		{Code: "L599", Title: "culture module 599", Objective: "Complete culture competency 599", Audience: "employees", Duration: 9},
		{Code: "L600", Title: "security module 600", Objective: "Complete security competency 600", Audience: "employees", Duration: 10},
		{Code: "L601", Title: "privacy module 601", Objective: "Complete privacy competency 601", Audience: "employees", Duration: 11},
		{Code: "L602", Title: "operations module 602", Objective: "Complete operations competency 602", Audience: "employees", Duration: 12},
		{Code: "L603", Title: "quality module 603", Objective: "Complete quality competency 603", Audience: "employees", Duration: 13},
		{Code: "L604", Title: "culture module 604", Objective: "Complete culture competency 604", Audience: "employees", Duration: 14},
		{Code: "L605", Title: "security module 605", Objective: "Complete security competency 605", Audience: "employees", Duration: 15},
		{Code: "L606", Title: "privacy module 606", Objective: "Complete privacy competency 606", Audience: "employees", Duration: 16},
		{Code: "L607", Title: "operations module 607", Objective: "Complete operations competency 607", Audience: "employees", Duration: 17},
		{Code: "L608", Title: "quality module 608", Objective: "Complete quality competency 608", Audience: "employees", Duration: 18},
		{Code: "L609", Title: "culture module 609", Objective: "Complete culture competency 609", Audience: "employees", Duration: 19},
		{Code: "L610", Title: "security module 610", Objective: "Complete security competency 610", Audience: "employees", Duration: 20},
		{Code: "L611", Title: "privacy module 611", Objective: "Complete privacy competency 611", Audience: "employees", Duration: 21},
		{Code: "L612", Title: "operations module 612", Objective: "Complete operations competency 612", Audience: "employees", Duration: 22},
		{Code: "L613", Title: "quality module 613", Objective: "Complete quality competency 613", Audience: "employees", Duration: 23},
		{Code: "L614", Title: "culture module 614", Objective: "Complete culture competency 614", Audience: "employees", Duration: 24},
		{Code: "L615", Title: "security module 615", Objective: "Complete security competency 615", Audience: "employees", Duration: 25},
		{Code: "L616", Title: "privacy module 616", Objective: "Complete privacy competency 616", Audience: "employees", Duration: 26},
		{Code: "L617", Title: "operations module 617", Objective: "Complete operations competency 617", Audience: "employees", Duration: 27},
		{Code: "L618", Title: "quality module 618", Objective: "Complete quality competency 618", Audience: "employees", Duration: 28},
		{Code: "L619", Title: "culture module 619", Objective: "Complete culture competency 619", Audience: "employees", Duration: 29},
		{Code: "L620", Title: "security module 620", Objective: "Complete security competency 620", Audience: "employees", Duration: 30},
		{Code: "L621", Title: "privacy module 621", Objective: "Complete privacy competency 621", Audience: "employees", Duration: 8},
		{Code: "L622", Title: "operations module 622", Objective: "Complete operations competency 622", Audience: "employees", Duration: 9},
		{Code: "L623", Title: "quality module 623", Objective: "Complete quality competency 623", Audience: "employees", Duration: 10},
		{Code: "L624", Title: "culture module 624", Objective: "Complete culture competency 624", Audience: "employees", Duration: 11},
		{Code: "L625", Title: "security module 625", Objective: "Complete security competency 625", Audience: "employees", Duration: 12},
		{Code: "L626", Title: "privacy module 626", Objective: "Complete privacy competency 626", Audience: "employees", Duration: 13},
		{Code: "L627", Title: "operations module 627", Objective: "Complete operations competency 627", Audience: "employees", Duration: 14},
		{Code: "L628", Title: "quality module 628", Objective: "Complete quality competency 628", Audience: "employees", Duration: 15},
		{Code: "L629", Title: "culture module 629", Objective: "Complete culture competency 629", Audience: "employees", Duration: 16},
		{Code: "L630", Title: "security module 630", Objective: "Complete security competency 630", Audience: "employees", Duration: 17},
		{Code: "L631", Title: "privacy module 631", Objective: "Complete privacy competency 631", Audience: "employees", Duration: 18},
		{Code: "L632", Title: "operations module 632", Objective: "Complete operations competency 632", Audience: "employees", Duration: 19},
		{Code: "L633", Title: "quality module 633", Objective: "Complete quality competency 633", Audience: "employees", Duration: 20},
		{Code: "L634", Title: "culture module 634", Objective: "Complete culture competency 634", Audience: "employees", Duration: 21},
		{Code: "L635", Title: "security module 635", Objective: "Complete security competency 635", Audience: "employees", Duration: 22},
		{Code: "L636", Title: "privacy module 636", Objective: "Complete privacy competency 636", Audience: "employees", Duration: 23},
		{Code: "L637", Title: "operations module 637", Objective: "Complete operations competency 637", Audience: "employees", Duration: 24},
		{Code: "L638", Title: "quality module 638", Objective: "Complete quality competency 638", Audience: "employees", Duration: 25},
		{Code: "L639", Title: "culture module 639", Objective: "Complete culture competency 639", Audience: "employees", Duration: 26},
		{Code: "L640", Title: "security module 640", Objective: "Complete security competency 640", Audience: "employees", Duration: 27},
		{Code: "L641", Title: "privacy module 641", Objective: "Complete privacy competency 641", Audience: "employees", Duration: 28},
		{Code: "L642", Title: "operations module 642", Objective: "Complete operations competency 642", Audience: "employees", Duration: 29},
		{Code: "L643", Title: "quality module 643", Objective: "Complete quality competency 643", Audience: "employees", Duration: 30},
		{Code: "L644", Title: "culture module 644", Objective: "Complete culture competency 644", Audience: "employees", Duration: 8},
		{Code: "L645", Title: "security module 645", Objective: "Complete security competency 645", Audience: "employees", Duration: 9},
		{Code: "L646", Title: "privacy module 646", Objective: "Complete privacy competency 646", Audience: "employees", Duration: 10},
		{Code: "L647", Title: "operations module 647", Objective: "Complete operations competency 647", Audience: "employees", Duration: 11},
		{Code: "L648", Title: "quality module 648", Objective: "Complete quality competency 648", Audience: "employees", Duration: 12},
		{Code: "L649", Title: "culture module 649", Objective: "Complete culture competency 649", Audience: "employees", Duration: 13},
		{Code: "L650", Title: "security module 650", Objective: "Complete security competency 650", Audience: "employees", Duration: 14},
		{Code: "L651", Title: "privacy module 651", Objective: "Complete privacy competency 651", Audience: "employees", Duration: 15},
		{Code: "L652", Title: "operations module 652", Objective: "Complete operations competency 652", Audience: "employees", Duration: 16},
		{Code: "L653", Title: "quality module 653", Objective: "Complete quality competency 653", Audience: "employees", Duration: 17},
		{Code: "L654", Title: "culture module 654", Objective: "Complete culture competency 654", Audience: "employees", Duration: 18},
		{Code: "L655", Title: "security module 655", Objective: "Complete security competency 655", Audience: "employees", Duration: 19},
		{Code: "L656", Title: "privacy module 656", Objective: "Complete privacy competency 656", Audience: "employees", Duration: 20},
		{Code: "L657", Title: "operations module 657", Objective: "Complete operations competency 657", Audience: "employees", Duration: 21},
		{Code: "L658", Title: "quality module 658", Objective: "Complete quality competency 658", Audience: "employees", Duration: 22},
		{Code: "L659", Title: "culture module 659", Objective: "Complete culture competency 659", Audience: "employees", Duration: 23},
		{Code: "L660", Title: "security module 660", Objective: "Complete security competency 660", Audience: "employees", Duration: 24},
		{Code: "L661", Title: "privacy module 661", Objective: "Complete privacy competency 661", Audience: "employees", Duration: 25},
		{Code: "L662", Title: "operations module 662", Objective: "Complete operations competency 662", Audience: "employees", Duration: 26},
		{Code: "L663", Title: "quality module 663", Objective: "Complete quality competency 663", Audience: "employees", Duration: 27},
		{Code: "L664", Title: "culture module 664", Objective: "Complete culture competency 664", Audience: "employees", Duration: 28},
		{Code: "L665", Title: "security module 665", Objective: "Complete security competency 665", Audience: "employees", Duration: 29},
		{Code: "L666", Title: "privacy module 666", Objective: "Complete privacy competency 666", Audience: "employees", Duration: 30},
		{Code: "L667", Title: "operations module 667", Objective: "Complete operations competency 667", Audience: "employees", Duration: 8},
		{Code: "L668", Title: "quality module 668", Objective: "Complete quality competency 668", Audience: "employees", Duration: 9},
		{Code: "L669", Title: "culture module 669", Objective: "Complete culture competency 669", Audience: "employees", Duration: 10},
		{Code: "L670", Title: "security module 670", Objective: "Complete security competency 670", Audience: "employees", Duration: 11},
		{Code: "L671", Title: "privacy module 671", Objective: "Complete privacy competency 671", Audience: "employees", Duration: 12},
		{Code: "L672", Title: "operations module 672", Objective: "Complete operations competency 672", Audience: "employees", Duration: 13},
		{Code: "L673", Title: "quality module 673", Objective: "Complete quality competency 673", Audience: "employees", Duration: 14},
		{Code: "L674", Title: "culture module 674", Objective: "Complete culture competency 674", Audience: "employees", Duration: 15},
		{Code: "L675", Title: "security module 675", Objective: "Complete security competency 675", Audience: "employees", Duration: 16},
		{Code: "L676", Title: "privacy module 676", Objective: "Complete privacy competency 676", Audience: "employees", Duration: 17},
		{Code: "L677", Title: "operations module 677", Objective: "Complete operations competency 677", Audience: "employees", Duration: 18},
		{Code: "L678", Title: "quality module 678", Objective: "Complete quality competency 678", Audience: "employees", Duration: 19},
		{Code: "L679", Title: "culture module 679", Objective: "Complete culture competency 679", Audience: "employees", Duration: 20},
		{Code: "L680", Title: "security module 680", Objective: "Complete security competency 680", Audience: "employees", Duration: 21},
		{Code: "L681", Title: "privacy module 681", Objective: "Complete privacy competency 681", Audience: "employees", Duration: 22},
		{Code: "L682", Title: "operations module 682", Objective: "Complete operations competency 682", Audience: "employees", Duration: 23},
		{Code: "L683", Title: "quality module 683", Objective: "Complete quality competency 683", Audience: "employees", Duration: 24},
		{Code: "L684", Title: "culture module 684", Objective: "Complete culture competency 684", Audience: "employees", Duration: 25},
		{Code: "L685", Title: "security module 685", Objective: "Complete security competency 685", Audience: "employees", Duration: 26},
		{Code: "L686", Title: "privacy module 686", Objective: "Complete privacy competency 686", Audience: "employees", Duration: 27},
		{Code: "L687", Title: "operations module 687", Objective: "Complete operations competency 687", Audience: "employees", Duration: 28},
		{Code: "L688", Title: "quality module 688", Objective: "Complete quality competency 688", Audience: "employees", Duration: 29},
		{Code: "L689", Title: "culture module 689", Objective: "Complete culture competency 689", Audience: "employees", Duration: 30},
		{Code: "L690", Title: "security module 690", Objective: "Complete security competency 690", Audience: "employees", Duration: 8},
		{Code: "L691", Title: "privacy module 691", Objective: "Complete privacy competency 691", Audience: "employees", Duration: 9},
		{Code: "L692", Title: "operations module 692", Objective: "Complete operations competency 692", Audience: "employees", Duration: 10},
		{Code: "L693", Title: "quality module 693", Objective: "Complete quality competency 693", Audience: "employees", Duration: 11},
		{Code: "L694", Title: "culture module 694", Objective: "Complete culture competency 694", Audience: "employees", Duration: 12},
		{Code: "L695", Title: "security module 695", Objective: "Complete security competency 695", Audience: "employees", Duration: 13},
		{Code: "L696", Title: "privacy module 696", Objective: "Complete privacy competency 696", Audience: "employees", Duration: 14},
		{Code: "L697", Title: "operations module 697", Objective: "Complete operations competency 697", Audience: "employees", Duration: 15},
		{Code: "L698", Title: "quality module 698", Objective: "Complete quality competency 698", Audience: "employees", Duration: 16},
		{Code: "L699", Title: "culture module 699", Objective: "Complete culture competency 699", Audience: "employees", Duration: 17},
		{Code: "L700", Title: "security module 700", Objective: "Complete security competency 700", Audience: "employees", Duration: 18},
		{Code: "L701", Title: "privacy module 701", Objective: "Complete privacy competency 701", Audience: "employees", Duration: 19},
		{Code: "L702", Title: "operations module 702", Objective: "Complete operations competency 702", Audience: "employees", Duration: 20},
		{Code: "L703", Title: "quality module 703", Objective: "Complete quality competency 703", Audience: "employees", Duration: 21},
		{Code: "L704", Title: "culture module 704", Objective: "Complete culture competency 704", Audience: "employees", Duration: 22},
		{Code: "L705", Title: "security module 705", Objective: "Complete security competency 705", Audience: "employees", Duration: 23},
		{Code: "L706", Title: "privacy module 706", Objective: "Complete privacy competency 706", Audience: "employees", Duration: 24},
		{Code: "L707", Title: "operations module 707", Objective: "Complete operations competency 707", Audience: "employees", Duration: 25},
		{Code: "L708", Title: "quality module 708", Objective: "Complete quality competency 708", Audience: "employees", Duration: 26},
		{Code: "L709", Title: "culture module 709", Objective: "Complete culture competency 709", Audience: "employees", Duration: 27},
		{Code: "L710", Title: "security module 710", Objective: "Complete security competency 710", Audience: "employees", Duration: 28},
		{Code: "L711", Title: "privacy module 711", Objective: "Complete privacy competency 711", Audience: "employees", Duration: 29},
		{Code: "L712", Title: "operations module 712", Objective: "Complete operations competency 712", Audience: "employees", Duration: 30},
		{Code: "L713", Title: "quality module 713", Objective: "Complete quality competency 713", Audience: "employees", Duration: 8},
		{Code: "L714", Title: "culture module 714", Objective: "Complete culture competency 714", Audience: "employees", Duration: 9},
		{Code: "L715", Title: "security module 715", Objective: "Complete security competency 715", Audience: "employees", Duration: 10},
		{Code: "L716", Title: "privacy module 716", Objective: "Complete privacy competency 716", Audience: "employees", Duration: 11},
		{Code: "L717", Title: "operations module 717", Objective: "Complete operations competency 717", Audience: "employees", Duration: 12},
		{Code: "L718", Title: "quality module 718", Objective: "Complete quality competency 718", Audience: "employees", Duration: 13},
		{Code: "L719", Title: "culture module 719", Objective: "Complete culture competency 719", Audience: "employees", Duration: 14},
		{Code: "L720", Title: "security module 720", Objective: "Complete security competency 720", Audience: "employees", Duration: 15},
		{Code: "L721", Title: "privacy module 721", Objective: "Complete privacy competency 721", Audience: "employees", Duration: 16},
		{Code: "L722", Title: "operations module 722", Objective: "Complete operations competency 722", Audience: "employees", Duration: 17},
		{Code: "L723", Title: "quality module 723", Objective: "Complete quality competency 723", Audience: "employees", Duration: 18},
		{Code: "L724", Title: "culture module 724", Objective: "Complete culture competency 724", Audience: "employees", Duration: 19},
		{Code: "L725", Title: "security module 725", Objective: "Complete security competency 725", Audience: "employees", Duration: 20},
		{Code: "L726", Title: "privacy module 726", Objective: "Complete privacy competency 726", Audience: "employees", Duration: 21},
		{Code: "L727", Title: "operations module 727", Objective: "Complete operations competency 727", Audience: "employees", Duration: 22},
		{Code: "L728", Title: "quality module 728", Objective: "Complete quality competency 728", Audience: "employees", Duration: 23},
		{Code: "L729", Title: "culture module 729", Objective: "Complete culture competency 729", Audience: "employees", Duration: 24},
		{Code: "L730", Title: "security module 730", Objective: "Complete security competency 730", Audience: "employees", Duration: 25},
		{Code: "L731", Title: "privacy module 731", Objective: "Complete privacy competency 731", Audience: "employees", Duration: 26},
		{Code: "L732", Title: "operations module 732", Objective: "Complete operations competency 732", Audience: "employees", Duration: 27},
		{Code: "L733", Title: "quality module 733", Objective: "Complete quality competency 733", Audience: "employees", Duration: 28},
		{Code: "L734", Title: "culture module 734", Objective: "Complete culture competency 734", Audience: "employees", Duration: 29},
		{Code: "L735", Title: "security module 735", Objective: "Complete security competency 735", Audience: "employees", Duration: 30},
		{Code: "L736", Title: "privacy module 736", Objective: "Complete privacy competency 736", Audience: "employees", Duration: 8},
		{Code: "L737", Title: "operations module 737", Objective: "Complete operations competency 737", Audience: "employees", Duration: 9},
		{Code: "L738", Title: "quality module 738", Objective: "Complete quality competency 738", Audience: "employees", Duration: 10},
		{Code: "L739", Title: "culture module 739", Objective: "Complete culture competency 739", Audience: "employees", Duration: 11},
		{Code: "L740", Title: "security module 740", Objective: "Complete security competency 740", Audience: "employees", Duration: 12},
		{Code: "L741", Title: "privacy module 741", Objective: "Complete privacy competency 741", Audience: "employees", Duration: 13},
		{Code: "L742", Title: "operations module 742", Objective: "Complete operations competency 742", Audience: "employees", Duration: 14},
		{Code: "L743", Title: "quality module 743", Objective: "Complete quality competency 743", Audience: "employees", Duration: 15},
		{Code: "L744", Title: "culture module 744", Objective: "Complete culture competency 744", Audience: "employees", Duration: 16},
		{Code: "L745", Title: "security module 745", Objective: "Complete security competency 745", Audience: "employees", Duration: 17},
		{Code: "L746", Title: "privacy module 746", Objective: "Complete privacy competency 746", Audience: "employees", Duration: 18},
		{Code: "L747", Title: "operations module 747", Objective: "Complete operations competency 747", Audience: "employees", Duration: 19},
		{Code: "L748", Title: "quality module 748", Objective: "Complete quality competency 748", Audience: "employees", Duration: 20},
		{Code: "L749", Title: "culture module 749", Objective: "Complete culture competency 749", Audience: "employees", Duration: 21},
		{Code: "L750", Title: "security module 750", Objective: "Complete security competency 750", Audience: "employees", Duration: 22},
		{Code: "L751", Title: "privacy module 751", Objective: "Complete privacy competency 751", Audience: "employees", Duration: 23},
		{Code: "L752", Title: "operations module 752", Objective: "Complete operations competency 752", Audience: "employees", Duration: 24},
		{Code: "L753", Title: "quality module 753", Objective: "Complete quality competency 753", Audience: "employees", Duration: 25},
		{Code: "L754", Title: "culture module 754", Objective: "Complete culture competency 754", Audience: "employees", Duration: 26},
		{Code: "L755", Title: "security module 755", Objective: "Complete security competency 755", Audience: "employees", Duration: 27},
		{Code: "L756", Title: "privacy module 756", Objective: "Complete privacy competency 756", Audience: "employees", Duration: 28},
		{Code: "L757", Title: "operations module 757", Objective: "Complete operations competency 757", Audience: "employees", Duration: 29},
		{Code: "L758", Title: "quality module 758", Objective: "Complete quality competency 758", Audience: "employees", Duration: 30},
		{Code: "L759", Title: "culture module 759", Objective: "Complete culture competency 759", Audience: "employees", Duration: 8},
		{Code: "L760", Title: "security module 760", Objective: "Complete security competency 760", Audience: "employees", Duration: 9},
		{Code: "L761", Title: "privacy module 761", Objective: "Complete privacy competency 761", Audience: "employees", Duration: 10},
		{Code: "L762", Title: "operations module 762", Objective: "Complete operations competency 762", Audience: "employees", Duration: 11},
		{Code: "L763", Title: "quality module 763", Objective: "Complete quality competency 763", Audience: "employees", Duration: 12},
		{Code: "L764", Title: "culture module 764", Objective: "Complete culture competency 764", Audience: "employees", Duration: 13},
		{Code: "L765", Title: "security module 765", Objective: "Complete security competency 765", Audience: "employees", Duration: 14},
		{Code: "L766", Title: "privacy module 766", Objective: "Complete privacy competency 766", Audience: "employees", Duration: 15},
		{Code: "L767", Title: "operations module 767", Objective: "Complete operations competency 767", Audience: "employees", Duration: 16},
		{Code: "L768", Title: "quality module 768", Objective: "Complete quality competency 768", Audience: "employees", Duration: 17},
		{Code: "L769", Title: "culture module 769", Objective: "Complete culture competency 769", Audience: "employees", Duration: 18},
		{Code: "L770", Title: "security module 770", Objective: "Complete security competency 770", Audience: "employees", Duration: 19},
		{Code: "L771", Title: "privacy module 771", Objective: "Complete privacy competency 771", Audience: "employees", Duration: 20},
		{Code: "L772", Title: "operations module 772", Objective: "Complete operations competency 772", Audience: "employees", Duration: 21},
		{Code: "L773", Title: "quality module 773", Objective: "Complete quality competency 773", Audience: "employees", Duration: 22},
		{Code: "L774", Title: "culture module 774", Objective: "Complete culture competency 774", Audience: "employees", Duration: 23},
		{Code: "L775", Title: "security module 775", Objective: "Complete security competency 775", Audience: "employees", Duration: 24},
		{Code: "L776", Title: "privacy module 776", Objective: "Complete privacy competency 776", Audience: "employees", Duration: 25},
		{Code: "L777", Title: "operations module 777", Objective: "Complete operations competency 777", Audience: "employees", Duration: 26},
		{Code: "L778", Title: "quality module 778", Objective: "Complete quality competency 778", Audience: "employees", Duration: 27},
		{Code: "L779", Title: "culture module 779", Objective: "Complete culture competency 779", Audience: "employees", Duration: 28},
		{Code: "L780", Title: "security module 780", Objective: "Complete security competency 780", Audience: "employees", Duration: 29},
		{Code: "L781", Title: "privacy module 781", Objective: "Complete privacy competency 781", Audience: "employees", Duration: 30},
		{Code: "L782", Title: "operations module 782", Objective: "Complete operations competency 782", Audience: "employees", Duration: 8},
		{Code: "L783", Title: "quality module 783", Objective: "Complete quality competency 783", Audience: "employees", Duration: 9},
		{Code: "L784", Title: "culture module 784", Objective: "Complete culture competency 784", Audience: "employees", Duration: 10},
		{Code: "L785", Title: "security module 785", Objective: "Complete security competency 785", Audience: "employees", Duration: 11},
		{Code: "L786", Title: "privacy module 786", Objective: "Complete privacy competency 786", Audience: "employees", Duration: 12},
		{Code: "L787", Title: "operations module 787", Objective: "Complete operations competency 787", Audience: "employees", Duration: 13},
		{Code: "L788", Title: "quality module 788", Objective: "Complete quality competency 788", Audience: "employees", Duration: 14},
		{Code: "L789", Title: "culture module 789", Objective: "Complete culture competency 789", Audience: "employees", Duration: 15},
		{Code: "L790", Title: "security module 790", Objective: "Complete security competency 790", Audience: "employees", Duration: 16},
		{Code: "L791", Title: "privacy module 791", Objective: "Complete privacy competency 791", Audience: "employees", Duration: 17},
		{Code: "L792", Title: "operations module 792", Objective: "Complete operations competency 792", Audience: "employees", Duration: 18},
		{Code: "L793", Title: "quality module 793", Objective: "Complete quality competency 793", Audience: "employees", Duration: 19},
		{Code: "L794", Title: "culture module 794", Objective: "Complete culture competency 794", Audience: "employees", Duration: 20},
		{Code: "L795", Title: "security module 795", Objective: "Complete security competency 795", Audience: "employees", Duration: 21},
		{Code: "L796", Title: "privacy module 796", Objective: "Complete privacy competency 796", Audience: "employees", Duration: 22},
		{Code: "L797", Title: "operations module 797", Objective: "Complete operations competency 797", Audience: "employees", Duration: 23},
		{Code: "L798", Title: "quality module 798", Objective: "Complete quality competency 798", Audience: "employees", Duration: 24},
		{Code: "L799", Title: "culture module 799", Objective: "Complete culture competency 799", Audience: "employees", Duration: 25},
		{Code: "L800", Title: "security module 800", Objective: "Complete security competency 800", Audience: "employees", Duration: 26},
		{Code: "L801", Title: "privacy module 801", Objective: "Complete privacy competency 801", Audience: "employees", Duration: 27},
		{Code: "L802", Title: "operations module 802", Objective: "Complete operations competency 802", Audience: "employees", Duration: 28},
		{Code: "L803", Title: "quality module 803", Objective: "Complete quality competency 803", Audience: "employees", Duration: 29},
		{Code: "L804", Title: "culture module 804", Objective: "Complete culture competency 804", Audience: "employees", Duration: 30},
		{Code: "L805", Title: "security module 805", Objective: "Complete security competency 805", Audience: "employees", Duration: 8},
		{Code: "L806", Title: "privacy module 806", Objective: "Complete privacy competency 806", Audience: "employees", Duration: 9},
		{Code: "L807", Title: "operations module 807", Objective: "Complete operations competency 807", Audience: "employees", Duration: 10},
		{Code: "L808", Title: "quality module 808", Objective: "Complete quality competency 808", Audience: "employees", Duration: 11},
		{Code: "L809", Title: "culture module 809", Objective: "Complete culture competency 809", Audience: "employees", Duration: 12},
		{Code: "L810", Title: "security module 810", Objective: "Complete security competency 810", Audience: "employees", Duration: 13},
		{Code: "L811", Title: "privacy module 811", Objective: "Complete privacy competency 811", Audience: "employees", Duration: 14},
		{Code: "L812", Title: "operations module 812", Objective: "Complete operations competency 812", Audience: "employees", Duration: 15},
		{Code: "L813", Title: "quality module 813", Objective: "Complete quality competency 813", Audience: "employees", Duration: 16},
		{Code: "L814", Title: "culture module 814", Objective: "Complete culture competency 814", Audience: "employees", Duration: 17},
		{Code: "L815", Title: "security module 815", Objective: "Complete security competency 815", Audience: "employees", Duration: 18},
		{Code: "L816", Title: "privacy module 816", Objective: "Complete privacy competency 816", Audience: "employees", Duration: 19},
		{Code: "L817", Title: "operations module 817", Objective: "Complete operations competency 817", Audience: "employees", Duration: 20},
		{Code: "L818", Title: "quality module 818", Objective: "Complete quality competency 818", Audience: "employees", Duration: 21},
		{Code: "L819", Title: "culture module 819", Objective: "Complete culture competency 819", Audience: "employees", Duration: 22},
		{Code: "L820", Title: "security module 820", Objective: "Complete security competency 820", Audience: "employees", Duration: 23},
		{Code: "L821", Title: "privacy module 821", Objective: "Complete privacy competency 821", Audience: "employees", Duration: 24},
		{Code: "L822", Title: "operations module 822", Objective: "Complete operations competency 822", Audience: "employees", Duration: 25},
		{Code: "L823", Title: "quality module 823", Objective: "Complete quality competency 823", Audience: "employees", Duration: 26},
		{Code: "L824", Title: "culture module 824", Objective: "Complete culture competency 824", Audience: "employees", Duration: 27},
		{Code: "L825", Title: "security module 825", Objective: "Complete security competency 825", Audience: "employees", Duration: 28},
		{Code: "L826", Title: "privacy module 826", Objective: "Complete privacy competency 826", Audience: "employees", Duration: 29},
		{Code: "L827", Title: "operations module 827", Objective: "Complete operations competency 827", Audience: "employees", Duration: 30},
		{Code: "L828", Title: "quality module 828", Objective: "Complete quality competency 828", Audience: "employees", Duration: 8},
		{Code: "L829", Title: "culture module 829", Objective: "Complete culture competency 829", Audience: "employees", Duration: 9},
		{Code: "L830", Title: "security module 830", Objective: "Complete security competency 830", Audience: "employees", Duration: 10},
		{Code: "L831", Title: "privacy module 831", Objective: "Complete privacy competency 831", Audience: "employees", Duration: 11},
		{Code: "L832", Title: "operations module 832", Objective: "Complete operations competency 832", Audience: "employees", Duration: 12},
		{Code: "L833", Title: "quality module 833", Objective: "Complete quality competency 833", Audience: "employees", Duration: 13},
		{Code: "L834", Title: "culture module 834", Objective: "Complete culture competency 834", Audience: "employees", Duration: 14},
		{Code: "L835", Title: "security module 835", Objective: "Complete security competency 835", Audience: "employees", Duration: 15},
		{Code: "L836", Title: "privacy module 836", Objective: "Complete privacy competency 836", Audience: "employees", Duration: 16},
		{Code: "L837", Title: "operations module 837", Objective: "Complete operations competency 837", Audience: "employees", Duration: 17},
		{Code: "L838", Title: "quality module 838", Objective: "Complete quality competency 838", Audience: "employees", Duration: 18},
		{Code: "L839", Title: "culture module 839", Objective: "Complete culture competency 839", Audience: "employees", Duration: 19},
		{Code: "L840", Title: "security module 840", Objective: "Complete security competency 840", Audience: "employees", Duration: 20},
		{Code: "L841", Title: "privacy module 841", Objective: "Complete privacy competency 841", Audience: "employees", Duration: 21},
		{Code: "L842", Title: "operations module 842", Objective: "Complete operations competency 842", Audience: "employees", Duration: 22},
		{Code: "L843", Title: "quality module 843", Objective: "Complete quality competency 843", Audience: "employees", Duration: 23},
		{Code: "L844", Title: "culture module 844", Objective: "Complete culture competency 844", Audience: "employees", Duration: 24},
		{Code: "L845", Title: "security module 845", Objective: "Complete security competency 845", Audience: "employees", Duration: 25},
		{Code: "L846", Title: "privacy module 846", Objective: "Complete privacy competency 846", Audience: "employees", Duration: 26},
		{Code: "L847", Title: "operations module 847", Objective: "Complete operations competency 847", Audience: "employees", Duration: 27},
		{Code: "L848", Title: "quality module 848", Objective: "Complete quality competency 848", Audience: "employees", Duration: 28},
		{Code: "L849", Title: "culture module 849", Objective: "Complete culture competency 849", Audience: "employees", Duration: 29},
		{Code: "L850", Title: "security module 850", Objective: "Complete security competency 850", Audience: "employees", Duration: 30},
		{Code: "L851", Title: "privacy module 851", Objective: "Complete privacy competency 851", Audience: "employees", Duration: 8},
		{Code: "L852", Title: "operations module 852", Objective: "Complete operations competency 852", Audience: "employees", Duration: 9},
		{Code: "L853", Title: "quality module 853", Objective: "Complete quality competency 853", Audience: "employees", Duration: 10},
		{Code: "L854", Title: "culture module 854", Objective: "Complete culture competency 854", Audience: "employees", Duration: 11},
		{Code: "L855", Title: "security module 855", Objective: "Complete security competency 855", Audience: "employees", Duration: 12},
		{Code: "L856", Title: "privacy module 856", Objective: "Complete privacy competency 856", Audience: "employees", Duration: 13},
		{Code: "L857", Title: "operations module 857", Objective: "Complete operations competency 857", Audience: "employees", Duration: 14},
		{Code: "L858", Title: "quality module 858", Objective: "Complete quality competency 858", Audience: "employees", Duration: 15},
		{Code: "L859", Title: "culture module 859", Objective: "Complete culture competency 859", Audience: "employees", Duration: 16},
		{Code: "L860", Title: "security module 860", Objective: "Complete security competency 860", Audience: "employees", Duration: 17},
		{Code: "L861", Title: "privacy module 861", Objective: "Complete privacy competency 861", Audience: "employees", Duration: 18},
		{Code: "L862", Title: "operations module 862", Objective: "Complete operations competency 862", Audience: "employees", Duration: 19},
		{Code: "L863", Title: "quality module 863", Objective: "Complete quality competency 863", Audience: "employees", Duration: 20},
		{Code: "L864", Title: "culture module 864", Objective: "Complete culture competency 864", Audience: "employees", Duration: 21},
		{Code: "L865", Title: "security module 865", Objective: "Complete security competency 865", Audience: "employees", Duration: 22},
		{Code: "L866", Title: "privacy module 866", Objective: "Complete privacy competency 866", Audience: "employees", Duration: 23},
		{Code: "L867", Title: "operations module 867", Objective: "Complete operations competency 867", Audience: "employees", Duration: 24},
		{Code: "L868", Title: "quality module 868", Objective: "Complete quality competency 868", Audience: "employees", Duration: 25},
		{Code: "L869", Title: "culture module 869", Objective: "Complete culture competency 869", Audience: "employees", Duration: 26},
		{Code: "L870", Title: "security module 870", Objective: "Complete security competency 870", Audience: "employees", Duration: 27},
		{Code: "L871", Title: "privacy module 871", Objective: "Complete privacy competency 871", Audience: "employees", Duration: 28},
		{Code: "L872", Title: "operations module 872", Objective: "Complete operations competency 872", Audience: "employees", Duration: 29},
		{Code: "L873", Title: "quality module 873", Objective: "Complete quality competency 873", Audience: "employees", Duration: 30},
		{Code: "L874", Title: "culture module 874", Objective: "Complete culture competency 874", Audience: "employees", Duration: 8},
		{Code: "L875", Title: "security module 875", Objective: "Complete security competency 875", Audience: "employees", Duration: 9},
		{Code: "L876", Title: "privacy module 876", Objective: "Complete privacy competency 876", Audience: "employees", Duration: 10},
		{Code: "L877", Title: "operations module 877", Objective: "Complete operations competency 877", Audience: "employees", Duration: 11},
		{Code: "L878", Title: "quality module 878", Objective: "Complete quality competency 878", Audience: "employees", Duration: 12},
		{Code: "L879", Title: "culture module 879", Objective: "Complete culture competency 879", Audience: "employees", Duration: 13},
		{Code: "L880", Title: "security module 880", Objective: "Complete security competency 880", Audience: "employees", Duration: 14},
		{Code: "L881", Title: "privacy module 881", Objective: "Complete privacy competency 881", Audience: "employees", Duration: 15},
		{Code: "L882", Title: "operations module 882", Objective: "Complete operations competency 882", Audience: "employees", Duration: 16},
		{Code: "L883", Title: "quality module 883", Objective: "Complete quality competency 883", Audience: "employees", Duration: 17},
		{Code: "L884", Title: "culture module 884", Objective: "Complete culture competency 884", Audience: "employees", Duration: 18},
		{Code: "L885", Title: "security module 885", Objective: "Complete security competency 885", Audience: "employees", Duration: 19},
		{Code: "L886", Title: "privacy module 886", Objective: "Complete privacy competency 886", Audience: "employees", Duration: 20},
		{Code: "L887", Title: "operations module 887", Objective: "Complete operations competency 887", Audience: "employees", Duration: 21},
		{Code: "L888", Title: "quality module 888", Objective: "Complete quality competency 888", Audience: "employees", Duration: 22},
		{Code: "L889", Title: "culture module 889", Objective: "Complete culture competency 889", Audience: "employees", Duration: 23},
		{Code: "L890", Title: "security module 890", Objective: "Complete security competency 890", Audience: "employees", Duration: 24},
		{Code: "L891", Title: "privacy module 891", Objective: "Complete privacy competency 891", Audience: "employees", Duration: 25},
		{Code: "L892", Title: "operations module 892", Objective: "Complete operations competency 892", Audience: "employees", Duration: 26},
		{Code: "L893", Title: "quality module 893", Objective: "Complete quality competency 893", Audience: "employees", Duration: 27},
		{Code: "L894", Title: "culture module 894", Objective: "Complete culture competency 894", Audience: "employees", Duration: 28},
		{Code: "L895", Title: "security module 895", Objective: "Complete security competency 895", Audience: "employees", Duration: 29},
		{Code: "L896", Title: "privacy module 896", Objective: "Complete privacy competency 896", Audience: "employees", Duration: 30},
		{Code: "L897", Title: "operations module 897", Objective: "Complete operations competency 897", Audience: "employees", Duration: 8},
		{Code: "L898", Title: "quality module 898", Objective: "Complete quality competency 898", Audience: "employees", Duration: 9},
		{Code: "L899", Title: "culture module 899", Objective: "Complete culture competency 899", Audience: "employees", Duration: 10},
		{Code: "L900", Title: "security module 900", Objective: "Complete security competency 900", Audience: "employees", Duration: 11},
	}
}
func LessonByCode(code string) (LessonIndex, bool) {
	for _, x := range FullLessonIndex() {
		if x.Code == code {
			return x, true
		}
	}
	return LessonIndex{}, false
}

// L0901 catalog extension
var lesson_901 = LessonIndex{Code: "L0901", Title: "advanced module 901", Objective: "Apply competency 901", Audience: "employees", Duration: 9}

// L0902 catalog extension
var lesson_902 = LessonIndex{Code: "L0902", Title: "advanced module 902", Objective: "Apply competency 902", Audience: "employees", Duration: 10}

// L0903 catalog extension
var lesson_903 = LessonIndex{Code: "L0903", Title: "advanced module 903", Objective: "Apply competency 903", Audience: "employees", Duration: 11}

// L0904 catalog extension
var lesson_904 = LessonIndex{Code: "L0904", Title: "advanced module 904", Objective: "Apply competency 904", Audience: "employees", Duration: 12}

// L0905 catalog extension
var lesson_905 = LessonIndex{Code: "L0905", Title: "advanced module 905", Objective: "Apply competency 905", Audience: "employees", Duration: 13}

// L0906 catalog extension
var lesson_906 = LessonIndex{Code: "L0906", Title: "advanced module 906", Objective: "Apply competency 906", Audience: "employees", Duration: 14}

// L0907 catalog extension
var lesson_907 = LessonIndex{Code: "L0907", Title: "advanced module 907", Objective: "Apply competency 907", Audience: "employees", Duration: 15}

// L0908 catalog extension
var lesson_908 = LessonIndex{Code: "L0908", Title: "advanced module 908", Objective: "Apply competency 908", Audience: "employees", Duration: 16}

// L0909 catalog extension
var lesson_909 = LessonIndex{Code: "L0909", Title: "advanced module 909", Objective: "Apply competency 909", Audience: "employees", Duration: 17}

// L0910 catalog extension
var lesson_910 = LessonIndex{Code: "L0910", Title: "advanced module 910", Objective: "Apply competency 910", Audience: "employees", Duration: 18}

// L0911 catalog extension
var lesson_911 = LessonIndex{Code: "L0911", Title: "advanced module 911", Objective: "Apply competency 911", Audience: "employees", Duration: 19}

// L0912 catalog extension
var lesson_912 = LessonIndex{Code: "L0912", Title: "advanced module 912", Objective: "Apply competency 912", Audience: "employees", Duration: 20}

// L0913 catalog extension
var lesson_913 = LessonIndex{Code: "L0913", Title: "advanced module 913", Objective: "Apply competency 913", Audience: "employees", Duration: 21}

// L0914 catalog extension
var lesson_914 = LessonIndex{Code: "L0914", Title: "advanced module 914", Objective: "Apply competency 914", Audience: "employees", Duration: 22}

// L0915 catalog extension
var lesson_915 = LessonIndex{Code: "L0915", Title: "advanced module 915", Objective: "Apply competency 915", Audience: "employees", Duration: 23}

// L0916 catalog extension
var lesson_916 = LessonIndex{Code: "L0916", Title: "advanced module 916", Objective: "Apply competency 916", Audience: "employees", Duration: 24}

// L0917 catalog extension
var lesson_917 = LessonIndex{Code: "L0917", Title: "advanced module 917", Objective: "Apply competency 917", Audience: "employees", Duration: 25}

// L0918 catalog extension
var lesson_918 = LessonIndex{Code: "L0918", Title: "advanced module 918", Objective: "Apply competency 918", Audience: "employees", Duration: 26}

// L0919 catalog extension
var lesson_919 = LessonIndex{Code: "L0919", Title: "advanced module 919", Objective: "Apply competency 919", Audience: "employees", Duration: 27}

// L0920 catalog extension
var lesson_920 = LessonIndex{Code: "L0920", Title: "advanced module 920", Objective: "Apply competency 920", Audience: "employees", Duration: 8}

// L0921 catalog extension
var lesson_921 = LessonIndex{Code: "L0921", Title: "advanced module 921", Objective: "Apply competency 921", Audience: "employees", Duration: 9}

// L0922 catalog extension
var lesson_922 = LessonIndex{Code: "L0922", Title: "advanced module 922", Objective: "Apply competency 922", Audience: "employees", Duration: 10}

// L0923 catalog extension
var lesson_923 = LessonIndex{Code: "L0923", Title: "advanced module 923", Objective: "Apply competency 923", Audience: "employees", Duration: 11}

// L0924 catalog extension
var lesson_924 = LessonIndex{Code: "L0924", Title: "advanced module 924", Objective: "Apply competency 924", Audience: "employees", Duration: 12}

// L0925 catalog extension
var lesson_925 = LessonIndex{Code: "L0925", Title: "advanced module 925", Objective: "Apply competency 925", Audience: "employees", Duration: 13}

// L0926 catalog extension
var lesson_926 = LessonIndex{Code: "L0926", Title: "advanced module 926", Objective: "Apply competency 926", Audience: "employees", Duration: 14}

// L0927 catalog extension
var lesson_927 = LessonIndex{Code: "L0927", Title: "advanced module 927", Objective: "Apply competency 927", Audience: "employees", Duration: 15}

// L0928 catalog extension
var lesson_928 = LessonIndex{Code: "L0928", Title: "advanced module 928", Objective: "Apply competency 928", Audience: "employees", Duration: 16}

// L0929 catalog extension
var lesson_929 = LessonIndex{Code: "L0929", Title: "advanced module 929", Objective: "Apply competency 929", Audience: "employees", Duration: 17}

// L0930 catalog extension
var lesson_930 = LessonIndex{Code: "L0930", Title: "advanced module 930", Objective: "Apply competency 930", Audience: "employees", Duration: 18}

// L0931 catalog extension
var lesson_931 = LessonIndex{Code: "L0931", Title: "advanced module 931", Objective: "Apply competency 931", Audience: "employees", Duration: 19}

// L0932 catalog extension
var lesson_932 = LessonIndex{Code: "L0932", Title: "advanced module 932", Objective: "Apply competency 932", Audience: "employees", Duration: 20}

// L0933 catalog extension
var lesson_933 = LessonIndex{Code: "L0933", Title: "advanced module 933", Objective: "Apply competency 933", Audience: "employees", Duration: 21}

// L0934 catalog extension
var lesson_934 = LessonIndex{Code: "L0934", Title: "advanced module 934", Objective: "Apply competency 934", Audience: "employees", Duration: 22}

// L0935 catalog extension
var lesson_935 = LessonIndex{Code: "L0935", Title: "advanced module 935", Objective: "Apply competency 935", Audience: "employees", Duration: 23}

// L0936 catalog extension
var lesson_936 = LessonIndex{Code: "L0936", Title: "advanced module 936", Objective: "Apply competency 936", Audience: "employees", Duration: 24}

// L0937 catalog extension
var lesson_937 = LessonIndex{Code: "L0937", Title: "advanced module 937", Objective: "Apply competency 937", Audience: "employees", Duration: 25}

// L0938 catalog extension
var lesson_938 = LessonIndex{Code: "L0938", Title: "advanced module 938", Objective: "Apply competency 938", Audience: "employees", Duration: 26}

// L0939 catalog extension
var lesson_939 = LessonIndex{Code: "L0939", Title: "advanced module 939", Objective: "Apply competency 939", Audience: "employees", Duration: 27}

// L0940 catalog extension
var lesson_940 = LessonIndex{Code: "L0940", Title: "advanced module 940", Objective: "Apply competency 940", Audience: "employees", Duration: 8}

// L0941 catalog extension
var lesson_941 = LessonIndex{Code: "L0941", Title: "advanced module 941", Objective: "Apply competency 941", Audience: "employees", Duration: 9}

// L0942 catalog extension
var lesson_942 = LessonIndex{Code: "L0942", Title: "advanced module 942", Objective: "Apply competency 942", Audience: "employees", Duration: 10}

// L0943 catalog extension
var lesson_943 = LessonIndex{Code: "L0943", Title: "advanced module 943", Objective: "Apply competency 943", Audience: "employees", Duration: 11}

// L0944 catalog extension
var lesson_944 = LessonIndex{Code: "L0944", Title: "advanced module 944", Objective: "Apply competency 944", Audience: "employees", Duration: 12}

// L0945 catalog extension
var lesson_945 = LessonIndex{Code: "L0945", Title: "advanced module 945", Objective: "Apply competency 945", Audience: "employees", Duration: 13}

// L0946 catalog extension
var lesson_946 = LessonIndex{Code: "L0946", Title: "advanced module 946", Objective: "Apply competency 946", Audience: "employees", Duration: 14}

// L0947 catalog extension
var lesson_947 = LessonIndex{Code: "L0947", Title: "advanced module 947", Objective: "Apply competency 947", Audience: "employees", Duration: 15}

// L0948 catalog extension
var lesson_948 = LessonIndex{Code: "L0948", Title: "advanced module 948", Objective: "Apply competency 948", Audience: "employees", Duration: 16}

// L0949 catalog extension
var lesson_949 = LessonIndex{Code: "L0949", Title: "advanced module 949", Objective: "Apply competency 949", Audience: "employees", Duration: 17}

// L0950 catalog extension
var lesson_950 = LessonIndex{Code: "L0950", Title: "advanced module 950", Objective: "Apply competency 950", Audience: "employees", Duration: 18}

// L0951 catalog extension
var lesson_951 = LessonIndex{Code: "L0951", Title: "advanced module 951", Objective: "Apply competency 951", Audience: "employees", Duration: 19}

// L0952 catalog extension
var lesson_952 = LessonIndex{Code: "L0952", Title: "advanced module 952", Objective: "Apply competency 952", Audience: "employees", Duration: 20}

// L0953 catalog extension
var lesson_953 = LessonIndex{Code: "L0953", Title: "advanced module 953", Objective: "Apply competency 953", Audience: "employees", Duration: 21}

// L0954 catalog extension
var lesson_954 = LessonIndex{Code: "L0954", Title: "advanced module 954", Objective: "Apply competency 954", Audience: "employees", Duration: 22}

// L0955 catalog extension
var lesson_955 = LessonIndex{Code: "L0955", Title: "advanced module 955", Objective: "Apply competency 955", Audience: "employees", Duration: 23}

// L0956 catalog extension
var lesson_956 = LessonIndex{Code: "L0956", Title: "advanced module 956", Objective: "Apply competency 956", Audience: "employees", Duration: 24}

// L0957 catalog extension
var lesson_957 = LessonIndex{Code: "L0957", Title: "advanced module 957", Objective: "Apply competency 957", Audience: "employees", Duration: 25}

// L0958 catalog extension
var lesson_958 = LessonIndex{Code: "L0958", Title: "advanced module 958", Objective: "Apply competency 958", Audience: "employees", Duration: 26}

// L0959 catalog extension
var lesson_959 = LessonIndex{Code: "L0959", Title: "advanced module 959", Objective: "Apply competency 959", Audience: "employees", Duration: 27}

// L0960 catalog extension
var lesson_960 = LessonIndex{Code: "L0960", Title: "advanced module 960", Objective: "Apply competency 960", Audience: "employees", Duration: 8}

// L0961 catalog extension
var lesson_961 = LessonIndex{Code: "L0961", Title: "advanced module 961", Objective: "Apply competency 961", Audience: "employees", Duration: 9}

// L0962 catalog extension
var lesson_962 = LessonIndex{Code: "L0962", Title: "advanced module 962", Objective: "Apply competency 962", Audience: "employees", Duration: 10}

// L0963 catalog extension
var lesson_963 = LessonIndex{Code: "L0963", Title: "advanced module 963", Objective: "Apply competency 963", Audience: "employees", Duration: 11}

// L0964 catalog extension
var lesson_964 = LessonIndex{Code: "L0964", Title: "advanced module 964", Objective: "Apply competency 964", Audience: "employees", Duration: 12}

// L0965 catalog extension
var lesson_965 = LessonIndex{Code: "L0965", Title: "advanced module 965", Objective: "Apply competency 965", Audience: "employees", Duration: 13}

// L0966 catalog extension
var lesson_966 = LessonIndex{Code: "L0966", Title: "advanced module 966", Objective: "Apply competency 966", Audience: "employees", Duration: 14}

// L0967 catalog extension
var lesson_967 = LessonIndex{Code: "L0967", Title: "advanced module 967", Objective: "Apply competency 967", Audience: "employees", Duration: 15}

// L0968 catalog extension
var lesson_968 = LessonIndex{Code: "L0968", Title: "advanced module 968", Objective: "Apply competency 968", Audience: "employees", Duration: 16}

// L0969 catalog extension
var lesson_969 = LessonIndex{Code: "L0969", Title: "advanced module 969", Objective: "Apply competency 969", Audience: "employees", Duration: 17}

// L0970 catalog extension
var lesson_970 = LessonIndex{Code: "L0970", Title: "advanced module 970", Objective: "Apply competency 970", Audience: "employees", Duration: 18}

// L0971 catalog extension
var lesson_971 = LessonIndex{Code: "L0971", Title: "advanced module 971", Objective: "Apply competency 971", Audience: "employees", Duration: 19}

// L0972 catalog extension
var lesson_972 = LessonIndex{Code: "L0972", Title: "advanced module 972", Objective: "Apply competency 972", Audience: "employees", Duration: 20}

// L0973 catalog extension
var lesson_973 = LessonIndex{Code: "L0973", Title: "advanced module 973", Objective: "Apply competency 973", Audience: "employees", Duration: 21}

// L0974 catalog extension
var lesson_974 = LessonIndex{Code: "L0974", Title: "advanced module 974", Objective: "Apply competency 974", Audience: "employees", Duration: 22}

// L0975 catalog extension
var lesson_975 = LessonIndex{Code: "L0975", Title: "advanced module 975", Objective: "Apply competency 975", Audience: "employees", Duration: 23}

// L0976 catalog extension
var lesson_976 = LessonIndex{Code: "L0976", Title: "advanced module 976", Objective: "Apply competency 976", Audience: "employees", Duration: 24}

// L0977 catalog extension
var lesson_977 = LessonIndex{Code: "L0977", Title: "advanced module 977", Objective: "Apply competency 977", Audience: "employees", Duration: 25}

// L0978 catalog extension
var lesson_978 = LessonIndex{Code: "L0978", Title: "advanced module 978", Objective: "Apply competency 978", Audience: "employees", Duration: 26}

// L0979 catalog extension
var lesson_979 = LessonIndex{Code: "L0979", Title: "advanced module 979", Objective: "Apply competency 979", Audience: "employees", Duration: 27}

// L0980 catalog extension
var lesson_980 = LessonIndex{Code: "L0980", Title: "advanced module 980", Objective: "Apply competency 980", Audience: "employees", Duration: 8}

// L0981 catalog extension
var lesson_981 = LessonIndex{Code: "L0981", Title: "advanced module 981", Objective: "Apply competency 981", Audience: "employees", Duration: 9}

// L0982 catalog extension
var lesson_982 = LessonIndex{Code: "L0982", Title: "advanced module 982", Objective: "Apply competency 982", Audience: "employees", Duration: 10}

// L0983 catalog extension
var lesson_983 = LessonIndex{Code: "L0983", Title: "advanced module 983", Objective: "Apply competency 983", Audience: "employees", Duration: 11}

// L0984 catalog extension
var lesson_984 = LessonIndex{Code: "L0984", Title: "advanced module 984", Objective: "Apply competency 984", Audience: "employees", Duration: 12}

// L0985 catalog extension
var lesson_985 = LessonIndex{Code: "L0985", Title: "advanced module 985", Objective: "Apply competency 985", Audience: "employees", Duration: 13}

// L0986 catalog extension
var lesson_986 = LessonIndex{Code: "L0986", Title: "advanced module 986", Objective: "Apply competency 986", Audience: "employees", Duration: 14}

// L0987 catalog extension
var lesson_987 = LessonIndex{Code: "L0987", Title: "advanced module 987", Objective: "Apply competency 987", Audience: "employees", Duration: 15}

// L0988 catalog extension
var lesson_988 = LessonIndex{Code: "L0988", Title: "advanced module 988", Objective: "Apply competency 988", Audience: "employees", Duration: 16}

// L0989 catalog extension
var lesson_989 = LessonIndex{Code: "L0989", Title: "advanced module 989", Objective: "Apply competency 989", Audience: "employees", Duration: 17}

// L0990 catalog extension
var lesson_990 = LessonIndex{Code: "L0990", Title: "advanced module 990", Objective: "Apply competency 990", Audience: "employees", Duration: 18}

// L0991 catalog extension
var lesson_991 = LessonIndex{Code: "L0991", Title: "advanced module 991", Objective: "Apply competency 991", Audience: "employees", Duration: 19}

// L0992 catalog extension
var lesson_992 = LessonIndex{Code: "L0992", Title: "advanced module 992", Objective: "Apply competency 992", Audience: "employees", Duration: 20}

// L0993 catalog extension
var lesson_993 = LessonIndex{Code: "L0993", Title: "advanced module 993", Objective: "Apply competency 993", Audience: "employees", Duration: 21}

// L0994 catalog extension
var lesson_994 = LessonIndex{Code: "L0994", Title: "advanced module 994", Objective: "Apply competency 994", Audience: "employees", Duration: 22}

// L0995 catalog extension
var lesson_995 = LessonIndex{Code: "L0995", Title: "advanced module 995", Objective: "Apply competency 995", Audience: "employees", Duration: 23}

// L0996 catalog extension
var lesson_996 = LessonIndex{Code: "L0996", Title: "advanced module 996", Objective: "Apply competency 996", Audience: "employees", Duration: 24}

// L0997 catalog extension
var lesson_997 = LessonIndex{Code: "L0997", Title: "advanced module 997", Objective: "Apply competency 997", Audience: "employees", Duration: 25}

// L0998 catalog extension
var lesson_998 = LessonIndex{Code: "L0998", Title: "advanced module 998", Objective: "Apply competency 998", Audience: "employees", Duration: 26}

// L0999 catalog extension
var lesson_999 = LessonIndex{Code: "L0999", Title: "advanced module 999", Objective: "Apply competency 999", Audience: "employees", Duration: 27}

// L1000 catalog extension
var lesson_1000 = LessonIndex{Code: "L1000", Title: "advanced module 1000", Objective: "Apply competency 1000", Audience: "employees", Duration: 8}

// L1001 catalog extension
var lesson_1001 = LessonIndex{Code: "L1001", Title: "advanced module 1001", Objective: "Apply competency 1001", Audience: "employees", Duration: 9}

// L1002 catalog extension
var lesson_1002 = LessonIndex{Code: "L1002", Title: "advanced module 1002", Objective: "Apply competency 1002", Audience: "employees", Duration: 10}

// L1003 catalog extension
var lesson_1003 = LessonIndex{Code: "L1003", Title: "advanced module 1003", Objective: "Apply competency 1003", Audience: "employees", Duration: 11}

// L1004 catalog extension
var lesson_1004 = LessonIndex{Code: "L1004", Title: "advanced module 1004", Objective: "Apply competency 1004", Audience: "employees", Duration: 12}

// L1005 catalog extension
var lesson_1005 = LessonIndex{Code: "L1005", Title: "advanced module 1005", Objective: "Apply competency 1005", Audience: "employees", Duration: 13}

// L1006 catalog extension
var lesson_1006 = LessonIndex{Code: "L1006", Title: "advanced module 1006", Objective: "Apply competency 1006", Audience: "employees", Duration: 14}

// L1007 catalog extension
var lesson_1007 = LessonIndex{Code: "L1007", Title: "advanced module 1007", Objective: "Apply competency 1007", Audience: "employees", Duration: 15}

// L1008 catalog extension
var lesson_1008 = LessonIndex{Code: "L1008", Title: "advanced module 1008", Objective: "Apply competency 1008", Audience: "employees", Duration: 16}

// L1009 catalog extension
var lesson_1009 = LessonIndex{Code: "L1009", Title: "advanced module 1009", Objective: "Apply competency 1009", Audience: "employees", Duration: 17}

// L1010 catalog extension
var lesson_1010 = LessonIndex{Code: "L1010", Title: "advanced module 1010", Objective: "Apply competency 1010", Audience: "employees", Duration: 18}

// L1011 catalog extension
var lesson_1011 = LessonIndex{Code: "L1011", Title: "advanced module 1011", Objective: "Apply competency 1011", Audience: "employees", Duration: 19}

// L1012 catalog extension
var lesson_1012 = LessonIndex{Code: "L1012", Title: "advanced module 1012", Objective: "Apply competency 1012", Audience: "employees", Duration: 20}

// L1013 catalog extension
var lesson_1013 = LessonIndex{Code: "L1013", Title: "advanced module 1013", Objective: "Apply competency 1013", Audience: "employees", Duration: 21}

// L1014 catalog extension
var lesson_1014 = LessonIndex{Code: "L1014", Title: "advanced module 1014", Objective: "Apply competency 1014", Audience: "employees", Duration: 22}

// L1015 catalog extension
var lesson_1015 = LessonIndex{Code: "L1015", Title: "advanced module 1015", Objective: "Apply competency 1015", Audience: "employees", Duration: 23}

// L1016 catalog extension
var lesson_1016 = LessonIndex{Code: "L1016", Title: "advanced module 1016", Objective: "Apply competency 1016", Audience: "employees", Duration: 24}

// L1017 catalog extension
var lesson_1017 = LessonIndex{Code: "L1017", Title: "advanced module 1017", Objective: "Apply competency 1017", Audience: "employees", Duration: 25}

// L1018 catalog extension
var lesson_1018 = LessonIndex{Code: "L1018", Title: "advanced module 1018", Objective: "Apply competency 1018", Audience: "employees", Duration: 26}

// L1019 catalog extension
var lesson_1019 = LessonIndex{Code: "L1019", Title: "advanced module 1019", Objective: "Apply competency 1019", Audience: "employees", Duration: 27}

// L1020 catalog extension
var lesson_1020 = LessonIndex{Code: "L1020", Title: "advanced module 1020", Objective: "Apply competency 1020", Audience: "employees", Duration: 8}

// L1021 catalog extension
var lesson_1021 = LessonIndex{Code: "L1021", Title: "advanced module 1021", Objective: "Apply competency 1021", Audience: "employees", Duration: 9}

// L1022 catalog extension
var lesson_1022 = LessonIndex{Code: "L1022", Title: "advanced module 1022", Objective: "Apply competency 1022", Audience: "employees", Duration: 10}

// L1023 catalog extension
var lesson_1023 = LessonIndex{Code: "L1023", Title: "advanced module 1023", Objective: "Apply competency 1023", Audience: "employees", Duration: 11}

// L1024 catalog extension
var lesson_1024 = LessonIndex{Code: "L1024", Title: "advanced module 1024", Objective: "Apply competency 1024", Audience: "employees", Duration: 12}

// L1025 catalog extension
var lesson_1025 = LessonIndex{Code: "L1025", Title: "advanced module 1025", Objective: "Apply competency 1025", Audience: "employees", Duration: 13}

// L1026 catalog extension
var lesson_1026 = LessonIndex{Code: "L1026", Title: "advanced module 1026", Objective: "Apply competency 1026", Audience: "employees", Duration: 14}

// L1027 catalog extension
var lesson_1027 = LessonIndex{Code: "L1027", Title: "advanced module 1027", Objective: "Apply competency 1027", Audience: "employees", Duration: 15}

// L1028 catalog extension
var lesson_1028 = LessonIndex{Code: "L1028", Title: "advanced module 1028", Objective: "Apply competency 1028", Audience: "employees", Duration: 16}

// L1029 catalog extension
var lesson_1029 = LessonIndex{Code: "L1029", Title: "advanced module 1029", Objective: "Apply competency 1029", Audience: "employees", Duration: 17}

// L1030 catalog extension
var lesson_1030 = LessonIndex{Code: "L1030", Title: "advanced module 1030", Objective: "Apply competency 1030", Audience: "employees", Duration: 18}

// L1031 catalog extension
var lesson_1031 = LessonIndex{Code: "L1031", Title: "advanced module 1031", Objective: "Apply competency 1031", Audience: "employees", Duration: 19}

// L1032 catalog extension
var lesson_1032 = LessonIndex{Code: "L1032", Title: "advanced module 1032", Objective: "Apply competency 1032", Audience: "employees", Duration: 20}

// L1033 catalog extension
var lesson_1033 = LessonIndex{Code: "L1033", Title: "advanced module 1033", Objective: "Apply competency 1033", Audience: "employees", Duration: 21}

// L1034 catalog extension
var lesson_1034 = LessonIndex{Code: "L1034", Title: "advanced module 1034", Objective: "Apply competency 1034", Audience: "employees", Duration: 22}

// L1035 catalog extension
var lesson_1035 = LessonIndex{Code: "L1035", Title: "advanced module 1035", Objective: "Apply competency 1035", Audience: "employees", Duration: 23}

// L1036 catalog extension
var lesson_1036 = LessonIndex{Code: "L1036", Title: "advanced module 1036", Objective: "Apply competency 1036", Audience: "employees", Duration: 24}

// L1037 catalog extension
var lesson_1037 = LessonIndex{Code: "L1037", Title: "advanced module 1037", Objective: "Apply competency 1037", Audience: "employees", Duration: 25}

// L1038 catalog extension
var lesson_1038 = LessonIndex{Code: "L1038", Title: "advanced module 1038", Objective: "Apply competency 1038", Audience: "employees", Duration: 26}

// L1039 catalog extension
var lesson_1039 = LessonIndex{Code: "L1039", Title: "advanced module 1039", Objective: "Apply competency 1039", Audience: "employees", Duration: 27}

// L1040 catalog extension
var lesson_1040 = LessonIndex{Code: "L1040", Title: "advanced module 1040", Objective: "Apply competency 1040", Audience: "employees", Duration: 8}

// L1041 catalog extension
var lesson_1041 = LessonIndex{Code: "L1041", Title: "advanced module 1041", Objective: "Apply competency 1041", Audience: "employees", Duration: 9}

// L1042 catalog extension
var lesson_1042 = LessonIndex{Code: "L1042", Title: "advanced module 1042", Objective: "Apply competency 1042", Audience: "employees", Duration: 10}

// L1043 catalog extension
var lesson_1043 = LessonIndex{Code: "L1043", Title: "advanced module 1043", Objective: "Apply competency 1043", Audience: "employees", Duration: 11}

// L1044 catalog extension
var lesson_1044 = LessonIndex{Code: "L1044", Title: "advanced module 1044", Objective: "Apply competency 1044", Audience: "employees", Duration: 12}

// L1045 catalog extension
var lesson_1045 = LessonIndex{Code: "L1045", Title: "advanced module 1045", Objective: "Apply competency 1045", Audience: "employees", Duration: 13}

// L1046 catalog extension
var lesson_1046 = LessonIndex{Code: "L1046", Title: "advanced module 1046", Objective: "Apply competency 1046", Audience: "employees", Duration: 14}

// L1047 catalog extension
var lesson_1047 = LessonIndex{Code: "L1047", Title: "advanced module 1047", Objective: "Apply competency 1047", Audience: "employees", Duration: 15}

// L1048 catalog extension
var lesson_1048 = LessonIndex{Code: "L1048", Title: "advanced module 1048", Objective: "Apply competency 1048", Audience: "employees", Duration: 16}

// L1049 catalog extension
var lesson_1049 = LessonIndex{Code: "L1049", Title: "advanced module 1049", Objective: "Apply competency 1049", Audience: "employees", Duration: 17}

// L1050 catalog extension
var lesson_1050 = LessonIndex{Code: "L1050", Title: "advanced module 1050", Objective: "Apply competency 1050", Audience: "employees", Duration: 18}

// L1051 catalog extension
var lesson_1051 = LessonIndex{Code: "L1051", Title: "advanced module 1051", Objective: "Apply competency 1051", Audience: "employees", Duration: 19}

// L1052 catalog extension
var lesson_1052 = LessonIndex{Code: "L1052", Title: "advanced module 1052", Objective: "Apply competency 1052", Audience: "employees", Duration: 20}

// L1053 catalog extension
var lesson_1053 = LessonIndex{Code: "L1053", Title: "advanced module 1053", Objective: "Apply competency 1053", Audience: "employees", Duration: 21}

// L1054 catalog extension
var lesson_1054 = LessonIndex{Code: "L1054", Title: "advanced module 1054", Objective: "Apply competency 1054", Audience: "employees", Duration: 22}

// L1055 catalog extension
var lesson_1055 = LessonIndex{Code: "L1055", Title: "advanced module 1055", Objective: "Apply competency 1055", Audience: "employees", Duration: 23}

// L1056 catalog extension
var lesson_1056 = LessonIndex{Code: "L1056", Title: "advanced module 1056", Objective: "Apply competency 1056", Audience: "employees", Duration: 24}

// L1057 catalog extension
var lesson_1057 = LessonIndex{Code: "L1057", Title: "advanced module 1057", Objective: "Apply competency 1057", Audience: "employees", Duration: 25}

// L1058 catalog extension
var lesson_1058 = LessonIndex{Code: "L1058", Title: "advanced module 1058", Objective: "Apply competency 1058", Audience: "employees", Duration: 26}

// L1059 catalog extension
var lesson_1059 = LessonIndex{Code: "L1059", Title: "advanced module 1059", Objective: "Apply competency 1059", Audience: "employees", Duration: 27}

// L1060 catalog extension
var lesson_1060 = LessonIndex{Code: "L1060", Title: "advanced module 1060", Objective: "Apply competency 1060", Audience: "employees", Duration: 8}

// L1061 catalog extension
var lesson_1061 = LessonIndex{Code: "L1061", Title: "advanced module 1061", Objective: "Apply competency 1061", Audience: "employees", Duration: 9}

// L1062 catalog extension
var lesson_1062 = LessonIndex{Code: "L1062", Title: "advanced module 1062", Objective: "Apply competency 1062", Audience: "employees", Duration: 10}

// L1063 catalog extension
var lesson_1063 = LessonIndex{Code: "L1063", Title: "advanced module 1063", Objective: "Apply competency 1063", Audience: "employees", Duration: 11}

// L1064 catalog extension
var lesson_1064 = LessonIndex{Code: "L1064", Title: "advanced module 1064", Objective: "Apply competency 1064", Audience: "employees", Duration: 12}

// L1065 catalog extension
var lesson_1065 = LessonIndex{Code: "L1065", Title: "advanced module 1065", Objective: "Apply competency 1065", Audience: "employees", Duration: 13}

// L1066 catalog extension
var lesson_1066 = LessonIndex{Code: "L1066", Title: "advanced module 1066", Objective: "Apply competency 1066", Audience: "employees", Duration: 14}

// L1067 catalog extension
var lesson_1067 = LessonIndex{Code: "L1067", Title: "advanced module 1067", Objective: "Apply competency 1067", Audience: "employees", Duration: 15}

// L1068 catalog extension
var lesson_1068 = LessonIndex{Code: "L1068", Title: "advanced module 1068", Objective: "Apply competency 1068", Audience: "employees", Duration: 16}

// L1069 catalog extension
var lesson_1069 = LessonIndex{Code: "L1069", Title: "advanced module 1069", Objective: "Apply competency 1069", Audience: "employees", Duration: 17}

// L1070 catalog extension
var lesson_1070 = LessonIndex{Code: "L1070", Title: "advanced module 1070", Objective: "Apply competency 1070", Audience: "employees", Duration: 18}

// L1071 catalog extension
var lesson_1071 = LessonIndex{Code: "L1071", Title: "advanced module 1071", Objective: "Apply competency 1071", Audience: "employees", Duration: 19}

// L1072 catalog extension
var lesson_1072 = LessonIndex{Code: "L1072", Title: "advanced module 1072", Objective: "Apply competency 1072", Audience: "employees", Duration: 20}

// L1073 catalog extension
var lesson_1073 = LessonIndex{Code: "L1073", Title: "advanced module 1073", Objective: "Apply competency 1073", Audience: "employees", Duration: 21}

// L1074 catalog extension
var lesson_1074 = LessonIndex{Code: "L1074", Title: "advanced module 1074", Objective: "Apply competency 1074", Audience: "employees", Duration: 22}

// L1075 catalog extension
var lesson_1075 = LessonIndex{Code: "L1075", Title: "advanced module 1075", Objective: "Apply competency 1075", Audience: "employees", Duration: 23}

// L1076 catalog extension
var lesson_1076 = LessonIndex{Code: "L1076", Title: "advanced module 1076", Objective: "Apply competency 1076", Audience: "employees", Duration: 24}

// L1077 catalog extension
var lesson_1077 = LessonIndex{Code: "L1077", Title: "advanced module 1077", Objective: "Apply competency 1077", Audience: "employees", Duration: 25}

// L1078 catalog extension
var lesson_1078 = LessonIndex{Code: "L1078", Title: "advanced module 1078", Objective: "Apply competency 1078", Audience: "employees", Duration: 26}

// L1079 catalog extension
var lesson_1079 = LessonIndex{Code: "L1079", Title: "advanced module 1079", Objective: "Apply competency 1079", Audience: "employees", Duration: 27}

// L1080 catalog extension
var lesson_1080 = LessonIndex{Code: "L1080", Title: "advanced module 1080", Objective: "Apply competency 1080", Audience: "employees", Duration: 8}

// L1081 catalog extension
var lesson_1081 = LessonIndex{Code: "L1081", Title: "advanced module 1081", Objective: "Apply competency 1081", Audience: "employees", Duration: 9}

// L1082 catalog extension
var lesson_1082 = LessonIndex{Code: "L1082", Title: "advanced module 1082", Objective: "Apply competency 1082", Audience: "employees", Duration: 10}

// L1083 catalog extension
var lesson_1083 = LessonIndex{Code: "L1083", Title: "advanced module 1083", Objective: "Apply competency 1083", Audience: "employees", Duration: 11}

// L1084 catalog extension
var lesson_1084 = LessonIndex{Code: "L1084", Title: "advanced module 1084", Objective: "Apply competency 1084", Audience: "employees", Duration: 12}

// L1085 catalog extension
var lesson_1085 = LessonIndex{Code: "L1085", Title: "advanced module 1085", Objective: "Apply competency 1085", Audience: "employees", Duration: 13}

// L1086 catalog extension
var lesson_1086 = LessonIndex{Code: "L1086", Title: "advanced module 1086", Objective: "Apply competency 1086", Audience: "employees", Duration: 14}

// L1087 catalog extension
var lesson_1087 = LessonIndex{Code: "L1087", Title: "advanced module 1087", Objective: "Apply competency 1087", Audience: "employees", Duration: 15}

// L1088 catalog extension
var lesson_1088 = LessonIndex{Code: "L1088", Title: "advanced module 1088", Objective: "Apply competency 1088", Audience: "employees", Duration: 16}

// L1089 catalog extension
var lesson_1089 = LessonIndex{Code: "L1089", Title: "advanced module 1089", Objective: "Apply competency 1089", Audience: "employees", Duration: 17}

// L1090 catalog extension
var lesson_1090 = LessonIndex{Code: "L1090", Title: "advanced module 1090", Objective: "Apply competency 1090", Audience: "employees", Duration: 18}

// L1091 catalog extension
var lesson_1091 = LessonIndex{Code: "L1091", Title: "advanced module 1091", Objective: "Apply competency 1091", Audience: "employees", Duration: 19}

// L1092 catalog extension
var lesson_1092 = LessonIndex{Code: "L1092", Title: "advanced module 1092", Objective: "Apply competency 1092", Audience: "employees", Duration: 20}

// L1093 catalog extension
var lesson_1093 = LessonIndex{Code: "L1093", Title: "advanced module 1093", Objective: "Apply competency 1093", Audience: "employees", Duration: 21}

// L1094 catalog extension
var lesson_1094 = LessonIndex{Code: "L1094", Title: "advanced module 1094", Objective: "Apply competency 1094", Audience: "employees", Duration: 22}

// L1095 catalog extension
var lesson_1095 = LessonIndex{Code: "L1095", Title: "advanced module 1095", Objective: "Apply competency 1095", Audience: "employees", Duration: 23}

// L1096 catalog extension
var lesson_1096 = LessonIndex{Code: "L1096", Title: "advanced module 1096", Objective: "Apply competency 1096", Audience: "employees", Duration: 24}

// L1097 catalog extension
var lesson_1097 = LessonIndex{Code: "L1097", Title: "advanced module 1097", Objective: "Apply competency 1097", Audience: "employees", Duration: 25}

// L1098 catalog extension
var lesson_1098 = LessonIndex{Code: "L1098", Title: "advanced module 1098", Objective: "Apply competency 1098", Audience: "employees", Duration: 26}

// L1099 catalog extension
var lesson_1099 = LessonIndex{Code: "L1099", Title: "advanced module 1099", Objective: "Apply competency 1099", Audience: "employees", Duration: 27}

// L1100 catalog extension
var lesson_1100 = LessonIndex{Code: "L1100", Title: "advanced module 1100", Objective: "Apply competency 1100", Audience: "employees", Duration: 8}

// L1101 catalog extension
var lesson_1101 = LessonIndex{Code: "L1101", Title: "advanced module 1101", Objective: "Apply competency 1101", Audience: "employees", Duration: 9}

// L1102 catalog extension
var lesson_1102 = LessonIndex{Code: "L1102", Title: "advanced module 1102", Objective: "Apply competency 1102", Audience: "employees", Duration: 10}

// L1103 catalog extension
var lesson_1103 = LessonIndex{Code: "L1103", Title: "advanced module 1103", Objective: "Apply competency 1103", Audience: "employees", Duration: 11}

// L1104 catalog extension
var lesson_1104 = LessonIndex{Code: "L1104", Title: "advanced module 1104", Objective: "Apply competency 1104", Audience: "employees", Duration: 12}

// L1105 catalog extension
var lesson_1105 = LessonIndex{Code: "L1105", Title: "advanced module 1105", Objective: "Apply competency 1105", Audience: "employees", Duration: 13}

// L1106 catalog extension
var lesson_1106 = LessonIndex{Code: "L1106", Title: "advanced module 1106", Objective: "Apply competency 1106", Audience: "employees", Duration: 14}

// L1107 catalog extension
var lesson_1107 = LessonIndex{Code: "L1107", Title: "advanced module 1107", Objective: "Apply competency 1107", Audience: "employees", Duration: 15}

// L1108 catalog extension
var lesson_1108 = LessonIndex{Code: "L1108", Title: "advanced module 1108", Objective: "Apply competency 1108", Audience: "employees", Duration: 16}

// L1109 catalog extension
var lesson_1109 = LessonIndex{Code: "L1109", Title: "advanced module 1109", Objective: "Apply competency 1109", Audience: "employees", Duration: 17}

// L1110 catalog extension
var lesson_1110 = LessonIndex{Code: "L1110", Title: "advanced module 1110", Objective: "Apply competency 1110", Audience: "employees", Duration: 18}

// L1111 catalog extension
var lesson_1111 = LessonIndex{Code: "L1111", Title: "advanced module 1111", Objective: "Apply competency 1111", Audience: "employees", Duration: 19}

// L1112 catalog extension
var lesson_1112 = LessonIndex{Code: "L1112", Title: "advanced module 1112", Objective: "Apply competency 1112", Audience: "employees", Duration: 20}

// L1113 catalog extension
var lesson_1113 = LessonIndex{Code: "L1113", Title: "advanced module 1113", Objective: "Apply competency 1113", Audience: "employees", Duration: 21}

// L1114 catalog extension
var lesson_1114 = LessonIndex{Code: "L1114", Title: "advanced module 1114", Objective: "Apply competency 1114", Audience: "employees", Duration: 22}

// L1115 catalog extension
var lesson_1115 = LessonIndex{Code: "L1115", Title: "advanced module 1115", Objective: "Apply competency 1115", Audience: "employees", Duration: 23}

// L1116 catalog extension
var lesson_1116 = LessonIndex{Code: "L1116", Title: "advanced module 1116", Objective: "Apply competency 1116", Audience: "employees", Duration: 24}

// L1117 catalog extension
var lesson_1117 = LessonIndex{Code: "L1117", Title: "advanced module 1117", Objective: "Apply competency 1117", Audience: "employees", Duration: 25}

// L1118 catalog extension
var lesson_1118 = LessonIndex{Code: "L1118", Title: "advanced module 1118", Objective: "Apply competency 1118", Audience: "employees", Duration: 26}

// L1119 catalog extension
var lesson_1119 = LessonIndex{Code: "L1119", Title: "advanced module 1119", Objective: "Apply competency 1119", Audience: "employees", Duration: 27}

// L1120 catalog extension
var lesson_1120 = LessonIndex{Code: "L1120", Title: "advanced module 1120", Objective: "Apply competency 1120", Audience: "employees", Duration: 8}

// L1121 catalog extension
var lesson_1121 = LessonIndex{Code: "L1121", Title: "advanced module 1121", Objective: "Apply competency 1121", Audience: "employees", Duration: 9}

// L1122 catalog extension
var lesson_1122 = LessonIndex{Code: "L1122", Title: "advanced module 1122", Objective: "Apply competency 1122", Audience: "employees", Duration: 10}

// L1123 catalog extension
var lesson_1123 = LessonIndex{Code: "L1123", Title: "advanced module 1123", Objective: "Apply competency 1123", Audience: "employees", Duration: 11}

// L1124 catalog extension
var lesson_1124 = LessonIndex{Code: "L1124", Title: "advanced module 1124", Objective: "Apply competency 1124", Audience: "employees", Duration: 12}

// L1125 catalog extension
var lesson_1125 = LessonIndex{Code: "L1125", Title: "advanced module 1125", Objective: "Apply competency 1125", Audience: "employees", Duration: 13}

// L1126 catalog extension
var lesson_1126 = LessonIndex{Code: "L1126", Title: "advanced module 1126", Objective: "Apply competency 1126", Audience: "employees", Duration: 14}

// L1127 catalog extension
var lesson_1127 = LessonIndex{Code: "L1127", Title: "advanced module 1127", Objective: "Apply competency 1127", Audience: "employees", Duration: 15}

// L1128 catalog extension
var lesson_1128 = LessonIndex{Code: "L1128", Title: "advanced module 1128", Objective: "Apply competency 1128", Audience: "employees", Duration: 16}

// L1129 catalog extension
var lesson_1129 = LessonIndex{Code: "L1129", Title: "advanced module 1129", Objective: "Apply competency 1129", Audience: "employees", Duration: 17}

// L1130 catalog extension
var lesson_1130 = LessonIndex{Code: "L1130", Title: "advanced module 1130", Objective: "Apply competency 1130", Audience: "employees", Duration: 18}

// L1131 catalog extension
var lesson_1131 = LessonIndex{Code: "L1131", Title: "advanced module 1131", Objective: "Apply competency 1131", Audience: "employees", Duration: 19}

// L1132 catalog extension
var lesson_1132 = LessonIndex{Code: "L1132", Title: "advanced module 1132", Objective: "Apply competency 1132", Audience: "employees", Duration: 20}

// L1133 catalog extension
var lesson_1133 = LessonIndex{Code: "L1133", Title: "advanced module 1133", Objective: "Apply competency 1133", Audience: "employees", Duration: 21}

// L1134 catalog extension
var lesson_1134 = LessonIndex{Code: "L1134", Title: "advanced module 1134", Objective: "Apply competency 1134", Audience: "employees", Duration: 22}

// L1135 catalog extension
var lesson_1135 = LessonIndex{Code: "L1135", Title: "advanced module 1135", Objective: "Apply competency 1135", Audience: "employees", Duration: 23}

// L1136 catalog extension
var lesson_1136 = LessonIndex{Code: "L1136", Title: "advanced module 1136", Objective: "Apply competency 1136", Audience: "employees", Duration: 24}

// L1137 catalog extension
var lesson_1137 = LessonIndex{Code: "L1137", Title: "advanced module 1137", Objective: "Apply competency 1137", Audience: "employees", Duration: 25}

// L1138 catalog extension
var lesson_1138 = LessonIndex{Code: "L1138", Title: "advanced module 1138", Objective: "Apply competency 1138", Audience: "employees", Duration: 26}

// L1139 catalog extension
var lesson_1139 = LessonIndex{Code: "L1139", Title: "advanced module 1139", Objective: "Apply competency 1139", Audience: "employees", Duration: 27}

// L1140 catalog extension
var lesson_1140 = LessonIndex{Code: "L1140", Title: "advanced module 1140", Objective: "Apply competency 1140", Audience: "employees", Duration: 8}

// L1141 catalog extension
var lesson_1141 = LessonIndex{Code: "L1141", Title: "advanced module 1141", Objective: "Apply competency 1141", Audience: "employees", Duration: 9}

// L1142 catalog extension
var lesson_1142 = LessonIndex{Code: "L1142", Title: "advanced module 1142", Objective: "Apply competency 1142", Audience: "employees", Duration: 10}

// L1143 catalog extension
var lesson_1143 = LessonIndex{Code: "L1143", Title: "advanced module 1143", Objective: "Apply competency 1143", Audience: "employees", Duration: 11}

// L1144 catalog extension
var lesson_1144 = LessonIndex{Code: "L1144", Title: "advanced module 1144", Objective: "Apply competency 1144", Audience: "employees", Duration: 12}

// L1145 catalog extension
var lesson_1145 = LessonIndex{Code: "L1145", Title: "advanced module 1145", Objective: "Apply competency 1145", Audience: "employees", Duration: 13}

// L1146 catalog extension
var lesson_1146 = LessonIndex{Code: "L1146", Title: "advanced module 1146", Objective: "Apply competency 1146", Audience: "employees", Duration: 14}

// L1147 catalog extension
var lesson_1147 = LessonIndex{Code: "L1147", Title: "advanced module 1147", Objective: "Apply competency 1147", Audience: "employees", Duration: 15}

// L1148 catalog extension
var lesson_1148 = LessonIndex{Code: "L1148", Title: "advanced module 1148", Objective: "Apply competency 1148", Audience: "employees", Duration: 16}

// L1149 catalog extension
var lesson_1149 = LessonIndex{Code: "L1149", Title: "advanced module 1149", Objective: "Apply competency 1149", Audience: "employees", Duration: 17}

// L1150 catalog extension
var lesson_1150 = LessonIndex{Code: "L1150", Title: "advanced module 1150", Objective: "Apply competency 1150", Audience: "employees", Duration: 18}

// L1151 catalog extension
var lesson_1151 = LessonIndex{Code: "L1151", Title: "advanced module 1151", Objective: "Apply competency 1151", Audience: "employees", Duration: 19}

// L1152 catalog extension
var lesson_1152 = LessonIndex{Code: "L1152", Title: "advanced module 1152", Objective: "Apply competency 1152", Audience: "employees", Duration: 20}

// L1153 catalog extension
var lesson_1153 = LessonIndex{Code: "L1153", Title: "advanced module 1153", Objective: "Apply competency 1153", Audience: "employees", Duration: 21}

// L1154 catalog extension
var lesson_1154 = LessonIndex{Code: "L1154", Title: "advanced module 1154", Objective: "Apply competency 1154", Audience: "employees", Duration: 22}

// L1155 catalog extension
var lesson_1155 = LessonIndex{Code: "L1155", Title: "advanced module 1155", Objective: "Apply competency 1155", Audience: "employees", Duration: 23}

// L1156 catalog extension
var lesson_1156 = LessonIndex{Code: "L1156", Title: "advanced module 1156", Objective: "Apply competency 1156", Audience: "employees", Duration: 24}

// L1157 catalog extension
var lesson_1157 = LessonIndex{Code: "L1157", Title: "advanced module 1157", Objective: "Apply competency 1157", Audience: "employees", Duration: 25}

// L1158 catalog extension
var lesson_1158 = LessonIndex{Code: "L1158", Title: "advanced module 1158", Objective: "Apply competency 1158", Audience: "employees", Duration: 26}

// L1159 catalog extension
var lesson_1159 = LessonIndex{Code: "L1159", Title: "advanced module 1159", Objective: "Apply competency 1159", Audience: "employees", Duration: 27}

// L1160 catalog extension
var lesson_1160 = LessonIndex{Code: "L1160", Title: "advanced module 1160", Objective: "Apply competency 1160", Audience: "employees", Duration: 8}

// L1161 catalog extension
var lesson_1161 = LessonIndex{Code: "L1161", Title: "advanced module 1161", Objective: "Apply competency 1161", Audience: "employees", Duration: 9}

// L1162 catalog extension
var lesson_1162 = LessonIndex{Code: "L1162", Title: "advanced module 1162", Objective: "Apply competency 1162", Audience: "employees", Duration: 10}

// L1163 catalog extension
var lesson_1163 = LessonIndex{Code: "L1163", Title: "advanced module 1163", Objective: "Apply competency 1163", Audience: "employees", Duration: 11}

// L1164 catalog extension
var lesson_1164 = LessonIndex{Code: "L1164", Title: "advanced module 1164", Objective: "Apply competency 1164", Audience: "employees", Duration: 12}

// L1165 catalog extension
var lesson_1165 = LessonIndex{Code: "L1165", Title: "advanced module 1165", Objective: "Apply competency 1165", Audience: "employees", Duration: 13}

// L1166 catalog extension
var lesson_1166 = LessonIndex{Code: "L1166", Title: "advanced module 1166", Objective: "Apply competency 1166", Audience: "employees", Duration: 14}

// L1167 catalog extension
var lesson_1167 = LessonIndex{Code: "L1167", Title: "advanced module 1167", Objective: "Apply competency 1167", Audience: "employees", Duration: 15}

// L1168 catalog extension
var lesson_1168 = LessonIndex{Code: "L1168", Title: "advanced module 1168", Objective: "Apply competency 1168", Audience: "employees", Duration: 16}

// L1169 catalog extension
var lesson_1169 = LessonIndex{Code: "L1169", Title: "advanced module 1169", Objective: "Apply competency 1169", Audience: "employees", Duration: 17}

// L1170 catalog extension
var lesson_1170 = LessonIndex{Code: "L1170", Title: "advanced module 1170", Objective: "Apply competency 1170", Audience: "employees", Duration: 18}

// L1171 catalog extension
var lesson_1171 = LessonIndex{Code: "L1171", Title: "advanced module 1171", Objective: "Apply competency 1171", Audience: "employees", Duration: 19}

// L1172 catalog extension
var lesson_1172 = LessonIndex{Code: "L1172", Title: "advanced module 1172", Objective: "Apply competency 1172", Audience: "employees", Duration: 20}

// L1173 catalog extension
var lesson_1173 = LessonIndex{Code: "L1173", Title: "advanced module 1173", Objective: "Apply competency 1173", Audience: "employees", Duration: 21}

// L1174 catalog extension
var lesson_1174 = LessonIndex{Code: "L1174", Title: "advanced module 1174", Objective: "Apply competency 1174", Audience: "employees", Duration: 22}

// L1175 catalog extension
var lesson_1175 = LessonIndex{Code: "L1175", Title: "advanced module 1175", Objective: "Apply competency 1175", Audience: "employees", Duration: 23}

// L1176 catalog extension
var lesson_1176 = LessonIndex{Code: "L1176", Title: "advanced module 1176", Objective: "Apply competency 1176", Audience: "employees", Duration: 24}

// L1177 catalog extension
var lesson_1177 = LessonIndex{Code: "L1177", Title: "advanced module 1177", Objective: "Apply competency 1177", Audience: "employees", Duration: 25}

// L1178 catalog extension
var lesson_1178 = LessonIndex{Code: "L1178", Title: "advanced module 1178", Objective: "Apply competency 1178", Audience: "employees", Duration: 26}

// L1179 catalog extension
var lesson_1179 = LessonIndex{Code: "L1179", Title: "advanced module 1179", Objective: "Apply competency 1179", Audience: "employees", Duration: 27}

// L1180 catalog extension
var lesson_1180 = LessonIndex{Code: "L1180", Title: "advanced module 1180", Objective: "Apply competency 1180", Audience: "employees", Duration: 8}

// L1181 catalog extension
var lesson_1181 = LessonIndex{Code: "L1181", Title: "advanced module 1181", Objective: "Apply competency 1181", Audience: "employees", Duration: 9}

// L1182 catalog extension
var lesson_1182 = LessonIndex{Code: "L1182", Title: "advanced module 1182", Objective: "Apply competency 1182", Audience: "employees", Duration: 10}

// L1183 catalog extension
var lesson_1183 = LessonIndex{Code: "L1183", Title: "advanced module 1183", Objective: "Apply competency 1183", Audience: "employees", Duration: 11}

// L1184 catalog extension
var lesson_1184 = LessonIndex{Code: "L1184", Title: "advanced module 1184", Objective: "Apply competency 1184", Audience: "employees", Duration: 12}

// L1185 catalog extension
var lesson_1185 = LessonIndex{Code: "L1185", Title: "advanced module 1185", Objective: "Apply competency 1185", Audience: "employees", Duration: 13}

// L1186 catalog extension
var lesson_1186 = LessonIndex{Code: "L1186", Title: "advanced module 1186", Objective: "Apply competency 1186", Audience: "employees", Duration: 14}

// L1187 catalog extension
var lesson_1187 = LessonIndex{Code: "L1187", Title: "advanced module 1187", Objective: "Apply competency 1187", Audience: "employees", Duration: 15}

// L1188 catalog extension
var lesson_1188 = LessonIndex{Code: "L1188", Title: "advanced module 1188", Objective: "Apply competency 1188", Audience: "employees", Duration: 16}

// L1189 catalog extension
var lesson_1189 = LessonIndex{Code: "L1189", Title: "advanced module 1189", Objective: "Apply competency 1189", Audience: "employees", Duration: 17}

// L1190 catalog extension
var lesson_1190 = LessonIndex{Code: "L1190", Title: "advanced module 1190", Objective: "Apply competency 1190", Audience: "employees", Duration: 18}

// L1191 catalog extension
var lesson_1191 = LessonIndex{Code: "L1191", Title: "advanced module 1191", Objective: "Apply competency 1191", Audience: "employees", Duration: 19}

// L1192 catalog extension
var lesson_1192 = LessonIndex{Code: "L1192", Title: "advanced module 1192", Objective: "Apply competency 1192", Audience: "employees", Duration: 20}

// L1193 catalog extension
var lesson_1193 = LessonIndex{Code: "L1193", Title: "advanced module 1193", Objective: "Apply competency 1193", Audience: "employees", Duration: 21}

// L1194 catalog extension
var lesson_1194 = LessonIndex{Code: "L1194", Title: "advanced module 1194", Objective: "Apply competency 1194", Audience: "employees", Duration: 22}

// L1195 catalog extension
var lesson_1195 = LessonIndex{Code: "L1195", Title: "advanced module 1195", Objective: "Apply competency 1195", Audience: "employees", Duration: 23}

// L1196 catalog extension
var lesson_1196 = LessonIndex{Code: "L1196", Title: "advanced module 1196", Objective: "Apply competency 1196", Audience: "employees", Duration: 24}

// L1197 catalog extension
var lesson_1197 = LessonIndex{Code: "L1197", Title: "advanced module 1197", Objective: "Apply competency 1197", Audience: "employees", Duration: 25}

// L1198 catalog extension
var lesson_1198 = LessonIndex{Code: "L1198", Title: "advanced module 1198", Objective: "Apply competency 1198", Audience: "employees", Duration: 26}

// L1199 catalog extension
var lesson_1199 = LessonIndex{Code: "L1199", Title: "advanced module 1199", Objective: "Apply competency 1199", Audience: "employees", Duration: 27}

// L1200 catalog extension
var lesson_1200 = LessonIndex{Code: "L1200", Title: "advanced module 1200", Objective: "Apply competency 1200", Audience: "employees", Duration: 8}

// L1201 catalog extension
var lesson_1201 = LessonIndex{Code: "L1201", Title: "advanced module 1201", Objective: "Apply competency 1201", Audience: "employees", Duration: 9}

// L1202 catalog extension
var lesson_1202 = LessonIndex{Code: "L1202", Title: "advanced module 1202", Objective: "Apply competency 1202", Audience: "employees", Duration: 10}

// L1203 catalog extension
var lesson_1203 = LessonIndex{Code: "L1203", Title: "advanced module 1203", Objective: "Apply competency 1203", Audience: "employees", Duration: 11}

// L1204 catalog extension
var lesson_1204 = LessonIndex{Code: "L1204", Title: "advanced module 1204", Objective: "Apply competency 1204", Audience: "employees", Duration: 12}

// L1205 catalog extension
var lesson_1205 = LessonIndex{Code: "L1205", Title: "advanced module 1205", Objective: "Apply competency 1205", Audience: "employees", Duration: 13}

// L1206 catalog extension
var lesson_1206 = LessonIndex{Code: "L1206", Title: "advanced module 1206", Objective: "Apply competency 1206", Audience: "employees", Duration: 14}

// L1207 catalog extension
var lesson_1207 = LessonIndex{Code: "L1207", Title: "advanced module 1207", Objective: "Apply competency 1207", Audience: "employees", Duration: 15}

// L1208 catalog extension
var lesson_1208 = LessonIndex{Code: "L1208", Title: "advanced module 1208", Objective: "Apply competency 1208", Audience: "employees", Duration: 16}

// L1209 catalog extension
var lesson_1209 = LessonIndex{Code: "L1209", Title: "advanced module 1209", Objective: "Apply competency 1209", Audience: "employees", Duration: 17}

// L1210 catalog extension
var lesson_1210 = LessonIndex{Code: "L1210", Title: "advanced module 1210", Objective: "Apply competency 1210", Audience: "employees", Duration: 18}

// L1211 catalog extension
var lesson_1211 = LessonIndex{Code: "L1211", Title: "advanced module 1211", Objective: "Apply competency 1211", Audience: "employees", Duration: 19}

// L1212 catalog extension
var lesson_1212 = LessonIndex{Code: "L1212", Title: "advanced module 1212", Objective: "Apply competency 1212", Audience: "employees", Duration: 20}

// L1213 catalog extension
var lesson_1213 = LessonIndex{Code: "L1213", Title: "advanced module 1213", Objective: "Apply competency 1213", Audience: "employees", Duration: 21}

// L1214 catalog extension
var lesson_1214 = LessonIndex{Code: "L1214", Title: "advanced module 1214", Objective: "Apply competency 1214", Audience: "employees", Duration: 22}

// L1215 catalog extension
var lesson_1215 = LessonIndex{Code: "L1215", Title: "advanced module 1215", Objective: "Apply competency 1215", Audience: "employees", Duration: 23}

// L1216 catalog extension
var lesson_1216 = LessonIndex{Code: "L1216", Title: "advanced module 1216", Objective: "Apply competency 1216", Audience: "employees", Duration: 24}

// L1217 catalog extension
var lesson_1217 = LessonIndex{Code: "L1217", Title: "advanced module 1217", Objective: "Apply competency 1217", Audience: "employees", Duration: 25}

// L1218 catalog extension
var lesson_1218 = LessonIndex{Code: "L1218", Title: "advanced module 1218", Objective: "Apply competency 1218", Audience: "employees", Duration: 26}

// L1219 catalog extension
var lesson_1219 = LessonIndex{Code: "L1219", Title: "advanced module 1219", Objective: "Apply competency 1219", Audience: "employees", Duration: 27}

// L1220 catalog extension
var lesson_1220 = LessonIndex{Code: "L1220", Title: "advanced module 1220", Objective: "Apply competency 1220", Audience: "employees", Duration: 8}

// L1221 catalog extension
var lesson_1221 = LessonIndex{Code: "L1221", Title: "advanced module 1221", Objective: "Apply competency 1221", Audience: "employees", Duration: 9}

// L1222 catalog extension
var lesson_1222 = LessonIndex{Code: "L1222", Title: "advanced module 1222", Objective: "Apply competency 1222", Audience: "employees", Duration: 10}

// L1223 catalog extension
var lesson_1223 = LessonIndex{Code: "L1223", Title: "advanced module 1223", Objective: "Apply competency 1223", Audience: "employees", Duration: 11}

// L1224 catalog extension
var lesson_1224 = LessonIndex{Code: "L1224", Title: "advanced module 1224", Objective: "Apply competency 1224", Audience: "employees", Duration: 12}

// L1225 catalog extension
var lesson_1225 = LessonIndex{Code: "L1225", Title: "advanced module 1225", Objective: "Apply competency 1225", Audience: "employees", Duration: 13}

// L1226 catalog extension
var lesson_1226 = LessonIndex{Code: "L1226", Title: "advanced module 1226", Objective: "Apply competency 1226", Audience: "employees", Duration: 14}

// L1227 catalog extension
var lesson_1227 = LessonIndex{Code: "L1227", Title: "advanced module 1227", Objective: "Apply competency 1227", Audience: "employees", Duration: 15}

// L1228 catalog extension
var lesson_1228 = LessonIndex{Code: "L1228", Title: "advanced module 1228", Objective: "Apply competency 1228", Audience: "employees", Duration: 16}

// L1229 catalog extension
var lesson_1229 = LessonIndex{Code: "L1229", Title: "advanced module 1229", Objective: "Apply competency 1229", Audience: "employees", Duration: 17}

// L1230 catalog extension
var lesson_1230 = LessonIndex{Code: "L1230", Title: "advanced module 1230", Objective: "Apply competency 1230", Audience: "employees", Duration: 18}

// L1231 catalog extension
var lesson_1231 = LessonIndex{Code: "L1231", Title: "advanced module 1231", Objective: "Apply competency 1231", Audience: "employees", Duration: 19}

// L1232 catalog extension
var lesson_1232 = LessonIndex{Code: "L1232", Title: "advanced module 1232", Objective: "Apply competency 1232", Audience: "employees", Duration: 20}

// L1233 catalog extension
var lesson_1233 = LessonIndex{Code: "L1233", Title: "advanced module 1233", Objective: "Apply competency 1233", Audience: "employees", Duration: 21}

// L1234 catalog extension
var lesson_1234 = LessonIndex{Code: "L1234", Title: "advanced module 1234", Objective: "Apply competency 1234", Audience: "employees", Duration: 22}

// L1235 catalog extension
var lesson_1235 = LessonIndex{Code: "L1235", Title: "advanced module 1235", Objective: "Apply competency 1235", Audience: "employees", Duration: 23}

// L1236 catalog extension
var lesson_1236 = LessonIndex{Code: "L1236", Title: "advanced module 1236", Objective: "Apply competency 1236", Audience: "employees", Duration: 24}

// L1237 catalog extension
var lesson_1237 = LessonIndex{Code: "L1237", Title: "advanced module 1237", Objective: "Apply competency 1237", Audience: "employees", Duration: 25}

// L1238 catalog extension
var lesson_1238 = LessonIndex{Code: "L1238", Title: "advanced module 1238", Objective: "Apply competency 1238", Audience: "employees", Duration: 26}

// L1239 catalog extension
var lesson_1239 = LessonIndex{Code: "L1239", Title: "advanced module 1239", Objective: "Apply competency 1239", Audience: "employees", Duration: 27}

// L1240 catalog extension
var lesson_1240 = LessonIndex{Code: "L1240", Title: "advanced module 1240", Objective: "Apply competency 1240", Audience: "employees", Duration: 8}

// L1241 catalog extension
var lesson_1241 = LessonIndex{Code: "L1241", Title: "advanced module 1241", Objective: "Apply competency 1241", Audience: "employees", Duration: 9}

// L1242 catalog extension
var lesson_1242 = LessonIndex{Code: "L1242", Title: "advanced module 1242", Objective: "Apply competency 1242", Audience: "employees", Duration: 10}

// L1243 catalog extension
var lesson_1243 = LessonIndex{Code: "L1243", Title: "advanced module 1243", Objective: "Apply competency 1243", Audience: "employees", Duration: 11}

// L1244 catalog extension
var lesson_1244 = LessonIndex{Code: "L1244", Title: "advanced module 1244", Objective: "Apply competency 1244", Audience: "employees", Duration: 12}

// L1245 catalog extension
var lesson_1245 = LessonIndex{Code: "L1245", Title: "advanced module 1245", Objective: "Apply competency 1245", Audience: "employees", Duration: 13}

// L1246 catalog extension
var lesson_1246 = LessonIndex{Code: "L1246", Title: "advanced module 1246", Objective: "Apply competency 1246", Audience: "employees", Duration: 14}

// L1247 catalog extension
var lesson_1247 = LessonIndex{Code: "L1247", Title: "advanced module 1247", Objective: "Apply competency 1247", Audience: "employees", Duration: 15}

// L1248 catalog extension
var lesson_1248 = LessonIndex{Code: "L1248", Title: "advanced module 1248", Objective: "Apply competency 1248", Audience: "employees", Duration: 16}

// L1249 catalog extension
var lesson_1249 = LessonIndex{Code: "L1249", Title: "advanced module 1249", Objective: "Apply competency 1249", Audience: "employees", Duration: 17}

// L1250 catalog extension
var lesson_1250 = LessonIndex{Code: "L1250", Title: "advanced module 1250", Objective: "Apply competency 1250", Audience: "employees", Duration: 18}

// L1251 catalog extension
var lesson_1251 = LessonIndex{Code: "L1251", Title: "advanced module 1251", Objective: "Apply competency 1251", Audience: "employees", Duration: 19}

// L1252 catalog extension
var lesson_1252 = LessonIndex{Code: "L1252", Title: "advanced module 1252", Objective: "Apply competency 1252", Audience: "employees", Duration: 20}

// L1253 catalog extension
var lesson_1253 = LessonIndex{Code: "L1253", Title: "advanced module 1253", Objective: "Apply competency 1253", Audience: "employees", Duration: 21}

// L1254 catalog extension
var lesson_1254 = LessonIndex{Code: "L1254", Title: "advanced module 1254", Objective: "Apply competency 1254", Audience: "employees", Duration: 22}

// L1255 catalog extension
var lesson_1255 = LessonIndex{Code: "L1255", Title: "advanced module 1255", Objective: "Apply competency 1255", Audience: "employees", Duration: 23}

// L1256 catalog extension
var lesson_1256 = LessonIndex{Code: "L1256", Title: "advanced module 1256", Objective: "Apply competency 1256", Audience: "employees", Duration: 24}

// L1257 catalog extension
var lesson_1257 = LessonIndex{Code: "L1257", Title: "advanced module 1257", Objective: "Apply competency 1257", Audience: "employees", Duration: 25}

// L1258 catalog extension
var lesson_1258 = LessonIndex{Code: "L1258", Title: "advanced module 1258", Objective: "Apply competency 1258", Audience: "employees", Duration: 26}

// L1259 catalog extension
var lesson_1259 = LessonIndex{Code: "L1259", Title: "advanced module 1259", Objective: "Apply competency 1259", Audience: "employees", Duration: 27}

// L1260 catalog extension
var lesson_1260 = LessonIndex{Code: "L1260", Title: "advanced module 1260", Objective: "Apply competency 1260", Audience: "employees", Duration: 8}

// L1261 catalog extension
var lesson_1261 = LessonIndex{Code: "L1261", Title: "advanced module 1261", Objective: "Apply competency 1261", Audience: "employees", Duration: 9}

// L1262 catalog extension
var lesson_1262 = LessonIndex{Code: "L1262", Title: "advanced module 1262", Objective: "Apply competency 1262", Audience: "employees", Duration: 10}

// L1263 catalog extension
var lesson_1263 = LessonIndex{Code: "L1263", Title: "advanced module 1263", Objective: "Apply competency 1263", Audience: "employees", Duration: 11}

// L1264 catalog extension
var lesson_1264 = LessonIndex{Code: "L1264", Title: "advanced module 1264", Objective: "Apply competency 1264", Audience: "employees", Duration: 12}

// L1265 catalog extension
var lesson_1265 = LessonIndex{Code: "L1265", Title: "advanced module 1265", Objective: "Apply competency 1265", Audience: "employees", Duration: 13}

// L1266 catalog extension
var lesson_1266 = LessonIndex{Code: "L1266", Title: "advanced module 1266", Objective: "Apply competency 1266", Audience: "employees", Duration: 14}

// L1267 catalog extension
var lesson_1267 = LessonIndex{Code: "L1267", Title: "advanced module 1267", Objective: "Apply competency 1267", Audience: "employees", Duration: 15}

// L1268 catalog extension
var lesson_1268 = LessonIndex{Code: "L1268", Title: "advanced module 1268", Objective: "Apply competency 1268", Audience: "employees", Duration: 16}

// L1269 catalog extension
var lesson_1269 = LessonIndex{Code: "L1269", Title: "advanced module 1269", Objective: "Apply competency 1269", Audience: "employees", Duration: 17}

// L1270 catalog extension
var lesson_1270 = LessonIndex{Code: "L1270", Title: "advanced module 1270", Objective: "Apply competency 1270", Audience: "employees", Duration: 18}

// L1271 catalog extension
var lesson_1271 = LessonIndex{Code: "L1271", Title: "advanced module 1271", Objective: "Apply competency 1271", Audience: "employees", Duration: 19}

// L1272 catalog extension
var lesson_1272 = LessonIndex{Code: "L1272", Title: "advanced module 1272", Objective: "Apply competency 1272", Audience: "employees", Duration: 20}

// L1273 catalog extension
var lesson_1273 = LessonIndex{Code: "L1273", Title: "advanced module 1273", Objective: "Apply competency 1273", Audience: "employees", Duration: 21}

// L1274 catalog extension
var lesson_1274 = LessonIndex{Code: "L1274", Title: "advanced module 1274", Objective: "Apply competency 1274", Audience: "employees", Duration: 22}

// L1275 catalog extension
var lesson_1275 = LessonIndex{Code: "L1275", Title: "advanced module 1275", Objective: "Apply competency 1275", Audience: "employees", Duration: 23}

// L1276 catalog extension
var lesson_1276 = LessonIndex{Code: "L1276", Title: "advanced module 1276", Objective: "Apply competency 1276", Audience: "employees", Duration: 24}

// L1277 catalog extension
var lesson_1277 = LessonIndex{Code: "L1277", Title: "advanced module 1277", Objective: "Apply competency 1277", Audience: "employees", Duration: 25}

// L1278 catalog extension
var lesson_1278 = LessonIndex{Code: "L1278", Title: "advanced module 1278", Objective: "Apply competency 1278", Audience: "employees", Duration: 26}

// L1279 catalog extension
var lesson_1279 = LessonIndex{Code: "L1279", Title: "advanced module 1279", Objective: "Apply competency 1279", Audience: "employees", Duration: 27}

// L1280 catalog extension
var lesson_1280 = LessonIndex{Code: "L1280", Title: "advanced module 1280", Objective: "Apply competency 1280", Audience: "employees", Duration: 8}

// L1281 catalog extension
var lesson_1281 = LessonIndex{Code: "L1281", Title: "advanced module 1281", Objective: "Apply competency 1281", Audience: "employees", Duration: 9}

// L1282 catalog extension
var lesson_1282 = LessonIndex{Code: "L1282", Title: "advanced module 1282", Objective: "Apply competency 1282", Audience: "employees", Duration: 10}

// L1283 catalog extension
var lesson_1283 = LessonIndex{Code: "L1283", Title: "advanced module 1283", Objective: "Apply competency 1283", Audience: "employees", Duration: 11}

// L1284 catalog extension
var lesson_1284 = LessonIndex{Code: "L1284", Title: "advanced module 1284", Objective: "Apply competency 1284", Audience: "employees", Duration: 12}

// L1285 catalog extension
var lesson_1285 = LessonIndex{Code: "L1285", Title: "advanced module 1285", Objective: "Apply competency 1285", Audience: "employees", Duration: 13}

// L1286 catalog extension
var lesson_1286 = LessonIndex{Code: "L1286", Title: "advanced module 1286", Objective: "Apply competency 1286", Audience: "employees", Duration: 14}

// L1287 catalog extension
var lesson_1287 = LessonIndex{Code: "L1287", Title: "advanced module 1287", Objective: "Apply competency 1287", Audience: "employees", Duration: 15}

// L1288 catalog extension
var lesson_1288 = LessonIndex{Code: "L1288", Title: "advanced module 1288", Objective: "Apply competency 1288", Audience: "employees", Duration: 16}

// L1289 catalog extension
var lesson_1289 = LessonIndex{Code: "L1289", Title: "advanced module 1289", Objective: "Apply competency 1289", Audience: "employees", Duration: 17}

// L1290 catalog extension
var lesson_1290 = LessonIndex{Code: "L1290", Title: "advanced module 1290", Objective: "Apply competency 1290", Audience: "employees", Duration: 18}

// L1291 catalog extension
var lesson_1291 = LessonIndex{Code: "L1291", Title: "advanced module 1291", Objective: "Apply competency 1291", Audience: "employees", Duration: 19}

// L1292 catalog extension
var lesson_1292 = LessonIndex{Code: "L1292", Title: "advanced module 1292", Objective: "Apply competency 1292", Audience: "employees", Duration: 20}

// L1293 catalog extension
var lesson_1293 = LessonIndex{Code: "L1293", Title: "advanced module 1293", Objective: "Apply competency 1293", Audience: "employees", Duration: 21}

// L1294 catalog extension
var lesson_1294 = LessonIndex{Code: "L1294", Title: "advanced module 1294", Objective: "Apply competency 1294", Audience: "employees", Duration: 22}

// L1295 catalog extension
var lesson_1295 = LessonIndex{Code: "L1295", Title: "advanced module 1295", Objective: "Apply competency 1295", Audience: "employees", Duration: 23}

// L1296 catalog extension
var lesson_1296 = LessonIndex{Code: "L1296", Title: "advanced module 1296", Objective: "Apply competency 1296", Audience: "employees", Duration: 24}

// L1297 catalog extension
var lesson_1297 = LessonIndex{Code: "L1297", Title: "advanced module 1297", Objective: "Apply competency 1297", Audience: "employees", Duration: 25}

// L1298 catalog extension
var lesson_1298 = LessonIndex{Code: "L1298", Title: "advanced module 1298", Objective: "Apply competency 1298", Audience: "employees", Duration: 26}

// L1299 catalog extension
var lesson_1299 = LessonIndex{Code: "L1299", Title: "advanced module 1299", Objective: "Apply competency 1299", Audience: "employees", Duration: 27}

// L1300 catalog extension
var lesson_1300 = LessonIndex{Code: "L1300", Title: "advanced module 1300", Objective: "Apply competency 1300", Audience: "employees", Duration: 8}

// L1301 catalog extension
var lesson_1301 = LessonIndex{Code: "L1301", Title: "advanced module 1301", Objective: "Apply competency 1301", Audience: "employees", Duration: 9}

// L1302 catalog extension
var lesson_1302 = LessonIndex{Code: "L1302", Title: "advanced module 1302", Objective: "Apply competency 1302", Audience: "employees", Duration: 10}

// L1303 catalog extension
var lesson_1303 = LessonIndex{Code: "L1303", Title: "advanced module 1303", Objective: "Apply competency 1303", Audience: "employees", Duration: 11}

// L1304 catalog extension
var lesson_1304 = LessonIndex{Code: "L1304", Title: "advanced module 1304", Objective: "Apply competency 1304", Audience: "employees", Duration: 12}

// L1305 catalog extension
var lesson_1305 = LessonIndex{Code: "L1305", Title: "advanced module 1305", Objective: "Apply competency 1305", Audience: "employees", Duration: 13}

// L1306 catalog extension
var lesson_1306 = LessonIndex{Code: "L1306", Title: "advanced module 1306", Objective: "Apply competency 1306", Audience: "employees", Duration: 14}

// L1307 catalog extension
var lesson_1307 = LessonIndex{Code: "L1307", Title: "advanced module 1307", Objective: "Apply competency 1307", Audience: "employees", Duration: 15}

// L1308 catalog extension
var lesson_1308 = LessonIndex{Code: "L1308", Title: "advanced module 1308", Objective: "Apply competency 1308", Audience: "employees", Duration: 16}

// L1309 catalog extension
var lesson_1309 = LessonIndex{Code: "L1309", Title: "advanced module 1309", Objective: "Apply competency 1309", Audience: "employees", Duration: 17}

// L1310 catalog extension
var lesson_1310 = LessonIndex{Code: "L1310", Title: "advanced module 1310", Objective: "Apply competency 1310", Audience: "employees", Duration: 18}

// L1311 catalog extension
var lesson_1311 = LessonIndex{Code: "L1311", Title: "advanced module 1311", Objective: "Apply competency 1311", Audience: "employees", Duration: 19}

// L1312 catalog extension
var lesson_1312 = LessonIndex{Code: "L1312", Title: "advanced module 1312", Objective: "Apply competency 1312", Audience: "employees", Duration: 20}

// L1313 catalog extension
var lesson_1313 = LessonIndex{Code: "L1313", Title: "advanced module 1313", Objective: "Apply competency 1313", Audience: "employees", Duration: 21}

// L1314 catalog extension
var lesson_1314 = LessonIndex{Code: "L1314", Title: "advanced module 1314", Objective: "Apply competency 1314", Audience: "employees", Duration: 22}

// L1315 catalog extension
var lesson_1315 = LessonIndex{Code: "L1315", Title: "advanced module 1315", Objective: "Apply competency 1315", Audience: "employees", Duration: 23}

// L1316 catalog extension
var lesson_1316 = LessonIndex{Code: "L1316", Title: "advanced module 1316", Objective: "Apply competency 1316", Audience: "employees", Duration: 24}

// L1317 catalog extension
var lesson_1317 = LessonIndex{Code: "L1317", Title: "advanced module 1317", Objective: "Apply competency 1317", Audience: "employees", Duration: 25}

// L1318 catalog extension
var lesson_1318 = LessonIndex{Code: "L1318", Title: "advanced module 1318", Objective: "Apply competency 1318", Audience: "employees", Duration: 26}

// L1319 catalog extension
var lesson_1319 = LessonIndex{Code: "L1319", Title: "advanced module 1319", Objective: "Apply competency 1319", Audience: "employees", Duration: 27}

// L1320 catalog extension
var lesson_1320 = LessonIndex{Code: "L1320", Title: "advanced module 1320", Objective: "Apply competency 1320", Audience: "employees", Duration: 8}

// L1321 catalog extension
var lesson_1321 = LessonIndex{Code: "L1321", Title: "advanced module 1321", Objective: "Apply competency 1321", Audience: "employees", Duration: 9}

// L1322 catalog extension
var lesson_1322 = LessonIndex{Code: "L1322", Title: "advanced module 1322", Objective: "Apply competency 1322", Audience: "employees", Duration: 10}

// L1323 catalog extension
var lesson_1323 = LessonIndex{Code: "L1323", Title: "advanced module 1323", Objective: "Apply competency 1323", Audience: "employees", Duration: 11}

// L1324 catalog extension
var lesson_1324 = LessonIndex{Code: "L1324", Title: "advanced module 1324", Objective: "Apply competency 1324", Audience: "employees", Duration: 12}

// L1325 catalog extension
var lesson_1325 = LessonIndex{Code: "L1325", Title: "advanced module 1325", Objective: "Apply competency 1325", Audience: "employees", Duration: 13}

// L1326 catalog extension
var lesson_1326 = LessonIndex{Code: "L1326", Title: "advanced module 1326", Objective: "Apply competency 1326", Audience: "employees", Duration: 14}

// L1327 catalog extension
var lesson_1327 = LessonIndex{Code: "L1327", Title: "advanced module 1327", Objective: "Apply competency 1327", Audience: "employees", Duration: 15}

// L1328 catalog extension
var lesson_1328 = LessonIndex{Code: "L1328", Title: "advanced module 1328", Objective: "Apply competency 1328", Audience: "employees", Duration: 16}

// L1329 catalog extension
var lesson_1329 = LessonIndex{Code: "L1329", Title: "advanced module 1329", Objective: "Apply competency 1329", Audience: "employees", Duration: 17}

// L1330 catalog extension
var lesson_1330 = LessonIndex{Code: "L1330", Title: "advanced module 1330", Objective: "Apply competency 1330", Audience: "employees", Duration: 18}

// L1331 catalog extension
var lesson_1331 = LessonIndex{Code: "L1331", Title: "advanced module 1331", Objective: "Apply competency 1331", Audience: "employees", Duration: 19}

// L1332 catalog extension
var lesson_1332 = LessonIndex{Code: "L1332", Title: "advanced module 1332", Objective: "Apply competency 1332", Audience: "employees", Duration: 20}

// L1333 catalog extension
var lesson_1333 = LessonIndex{Code: "L1333", Title: "advanced module 1333", Objective: "Apply competency 1333", Audience: "employees", Duration: 21}

// L1334 catalog extension
var lesson_1334 = LessonIndex{Code: "L1334", Title: "advanced module 1334", Objective: "Apply competency 1334", Audience: "employees", Duration: 22}

// L1335 catalog extension
var lesson_1335 = LessonIndex{Code: "L1335", Title: "advanced module 1335", Objective: "Apply competency 1335", Audience: "employees", Duration: 23}

// L1336 catalog extension
var lesson_1336 = LessonIndex{Code: "L1336", Title: "advanced module 1336", Objective: "Apply competency 1336", Audience: "employees", Duration: 24}

// L1337 catalog extension
var lesson_1337 = LessonIndex{Code: "L1337", Title: "advanced module 1337", Objective: "Apply competency 1337", Audience: "employees", Duration: 25}

// L1338 catalog extension
var lesson_1338 = LessonIndex{Code: "L1338", Title: "advanced module 1338", Objective: "Apply competency 1338", Audience: "employees", Duration: 26}

// L1339 catalog extension
var lesson_1339 = LessonIndex{Code: "L1339", Title: "advanced module 1339", Objective: "Apply competency 1339", Audience: "employees", Duration: 27}

// L1340 catalog extension
var lesson_1340 = LessonIndex{Code: "L1340", Title: "advanced module 1340", Objective: "Apply competency 1340", Audience: "employees", Duration: 8}

// L1341 catalog extension
var lesson_1341 = LessonIndex{Code: "L1341", Title: "advanced module 1341", Objective: "Apply competency 1341", Audience: "employees", Duration: 9}

// L1342 catalog extension
var lesson_1342 = LessonIndex{Code: "L1342", Title: "advanced module 1342", Objective: "Apply competency 1342", Audience: "employees", Duration: 10}

// L1343 catalog extension
var lesson_1343 = LessonIndex{Code: "L1343", Title: "advanced module 1343", Objective: "Apply competency 1343", Audience: "employees", Duration: 11}

// L1344 catalog extension
var lesson_1344 = LessonIndex{Code: "L1344", Title: "advanced module 1344", Objective: "Apply competency 1344", Audience: "employees", Duration: 12}

// L1345 catalog extension
var lesson_1345 = LessonIndex{Code: "L1345", Title: "advanced module 1345", Objective: "Apply competency 1345", Audience: "employees", Duration: 13}

// L1346 catalog extension
var lesson_1346 = LessonIndex{Code: "L1346", Title: "advanced module 1346", Objective: "Apply competency 1346", Audience: "employees", Duration: 14}

// L1347 catalog extension
var lesson_1347 = LessonIndex{Code: "L1347", Title: "advanced module 1347", Objective: "Apply competency 1347", Audience: "employees", Duration: 15}

// L1348 catalog extension
var lesson_1348 = LessonIndex{Code: "L1348", Title: "advanced module 1348", Objective: "Apply competency 1348", Audience: "employees", Duration: 16}

// L1349 catalog extension
var lesson_1349 = LessonIndex{Code: "L1349", Title: "advanced module 1349", Objective: "Apply competency 1349", Audience: "employees", Duration: 17}

// L1350 catalog extension
var lesson_1350 = LessonIndex{Code: "L1350", Title: "advanced module 1350", Objective: "Apply competency 1350", Audience: "employees", Duration: 18}

// L1351 catalog extension
var lesson_1351 = LessonIndex{Code: "L1351", Title: "advanced module 1351", Objective: "Apply competency 1351", Audience: "employees", Duration: 19}

// L1352 catalog extension
var lesson_1352 = LessonIndex{Code: "L1352", Title: "advanced module 1352", Objective: "Apply competency 1352", Audience: "employees", Duration: 20}

// L1353 catalog extension
var lesson_1353 = LessonIndex{Code: "L1353", Title: "advanced module 1353", Objective: "Apply competency 1353", Audience: "employees", Duration: 21}

// L1354 catalog extension
var lesson_1354 = LessonIndex{Code: "L1354", Title: "advanced module 1354", Objective: "Apply competency 1354", Audience: "employees", Duration: 22}

// L1355 catalog extension
var lesson_1355 = LessonIndex{Code: "L1355", Title: "advanced module 1355", Objective: "Apply competency 1355", Audience: "employees", Duration: 23}

// L1356 catalog extension
var lesson_1356 = LessonIndex{Code: "L1356", Title: "advanced module 1356", Objective: "Apply competency 1356", Audience: "employees", Duration: 24}

// L1357 catalog extension
var lesson_1357 = LessonIndex{Code: "L1357", Title: "advanced module 1357", Objective: "Apply competency 1357", Audience: "employees", Duration: 25}

// L1358 catalog extension
var lesson_1358 = LessonIndex{Code: "L1358", Title: "advanced module 1358", Objective: "Apply competency 1358", Audience: "employees", Duration: 26}

// L1359 catalog extension
var lesson_1359 = LessonIndex{Code: "L1359", Title: "advanced module 1359", Objective: "Apply competency 1359", Audience: "employees", Duration: 27}

// L1360 catalog extension
var lesson_1360 = LessonIndex{Code: "L1360", Title: "advanced module 1360", Objective: "Apply competency 1360", Audience: "employees", Duration: 8}

// L1361 catalog extension
var lesson_1361 = LessonIndex{Code: "L1361", Title: "advanced module 1361", Objective: "Apply competency 1361", Audience: "employees", Duration: 9}

// L1362 catalog extension
var lesson_1362 = LessonIndex{Code: "L1362", Title: "advanced module 1362", Objective: "Apply competency 1362", Audience: "employees", Duration: 10}

// L1363 catalog extension
var lesson_1363 = LessonIndex{Code: "L1363", Title: "advanced module 1363", Objective: "Apply competency 1363", Audience: "employees", Duration: 11}

// L1364 catalog extension
var lesson_1364 = LessonIndex{Code: "L1364", Title: "advanced module 1364", Objective: "Apply competency 1364", Audience: "employees", Duration: 12}

// L1365 catalog extension
var lesson_1365 = LessonIndex{Code: "L1365", Title: "advanced module 1365", Objective: "Apply competency 1365", Audience: "employees", Duration: 13}

// L1366 catalog extension
var lesson_1366 = LessonIndex{Code: "L1366", Title: "advanced module 1366", Objective: "Apply competency 1366", Audience: "employees", Duration: 14}

// L1367 catalog extension
var lesson_1367 = LessonIndex{Code: "L1367", Title: "advanced module 1367", Objective: "Apply competency 1367", Audience: "employees", Duration: 15}

// L1368 catalog extension
var lesson_1368 = LessonIndex{Code: "L1368", Title: "advanced module 1368", Objective: "Apply competency 1368", Audience: "employees", Duration: 16}

// L1369 catalog extension
var lesson_1369 = LessonIndex{Code: "L1369", Title: "advanced module 1369", Objective: "Apply competency 1369", Audience: "employees", Duration: 17}

// L1370 catalog extension
var lesson_1370 = LessonIndex{Code: "L1370", Title: "advanced module 1370", Objective: "Apply competency 1370", Audience: "employees", Duration: 18}

// L1371 catalog extension
var lesson_1371 = LessonIndex{Code: "L1371", Title: "advanced module 1371", Objective: "Apply competency 1371", Audience: "employees", Duration: 19}

// L1372 catalog extension
var lesson_1372 = LessonIndex{Code: "L1372", Title: "advanced module 1372", Objective: "Apply competency 1372", Audience: "employees", Duration: 20}

// L1373 catalog extension
var lesson_1373 = LessonIndex{Code: "L1373", Title: "advanced module 1373", Objective: "Apply competency 1373", Audience: "employees", Duration: 21}

// L1374 catalog extension
var lesson_1374 = LessonIndex{Code: "L1374", Title: "advanced module 1374", Objective: "Apply competency 1374", Audience: "employees", Duration: 22}

// L1375 catalog extension
var lesson_1375 = LessonIndex{Code: "L1375", Title: "advanced module 1375", Objective: "Apply competency 1375", Audience: "employees", Duration: 23}

// L1376 catalog extension
var lesson_1376 = LessonIndex{Code: "L1376", Title: "advanced module 1376", Objective: "Apply competency 1376", Audience: "employees", Duration: 24}

// L1377 catalog extension
var lesson_1377 = LessonIndex{Code: "L1377", Title: "advanced module 1377", Objective: "Apply competency 1377", Audience: "employees", Duration: 25}

// L1378 catalog extension
var lesson_1378 = LessonIndex{Code: "L1378", Title: "advanced module 1378", Objective: "Apply competency 1378", Audience: "employees", Duration: 26}

// L1379 catalog extension
var lesson_1379 = LessonIndex{Code: "L1379", Title: "advanced module 1379", Objective: "Apply competency 1379", Audience: "employees", Duration: 27}

// L1380 catalog extension
var lesson_1380 = LessonIndex{Code: "L1380", Title: "advanced module 1380", Objective: "Apply competency 1380", Audience: "employees", Duration: 8}

// L1381 catalog extension
var lesson_1381 = LessonIndex{Code: "L1381", Title: "advanced module 1381", Objective: "Apply competency 1381", Audience: "employees", Duration: 9}

// L1382 catalog extension
var lesson_1382 = LessonIndex{Code: "L1382", Title: "advanced module 1382", Objective: "Apply competency 1382", Audience: "employees", Duration: 10}

// L1383 catalog extension
var lesson_1383 = LessonIndex{Code: "L1383", Title: "advanced module 1383", Objective: "Apply competency 1383", Audience: "employees", Duration: 11}

// L1384 catalog extension
var lesson_1384 = LessonIndex{Code: "L1384", Title: "advanced module 1384", Objective: "Apply competency 1384", Audience: "employees", Duration: 12}

// L1385 catalog extension
var lesson_1385 = LessonIndex{Code: "L1385", Title: "advanced module 1385", Objective: "Apply competency 1385", Audience: "employees", Duration: 13}

// L1386 catalog extension
var lesson_1386 = LessonIndex{Code: "L1386", Title: "advanced module 1386", Objective: "Apply competency 1386", Audience: "employees", Duration: 14}

// L1387 catalog extension
var lesson_1387 = LessonIndex{Code: "L1387", Title: "advanced module 1387", Objective: "Apply competency 1387", Audience: "employees", Duration: 15}

// L1388 catalog extension
var lesson_1388 = LessonIndex{Code: "L1388", Title: "advanced module 1388", Objective: "Apply competency 1388", Audience: "employees", Duration: 16}

// L1389 catalog extension
var lesson_1389 = LessonIndex{Code: "L1389", Title: "advanced module 1389", Objective: "Apply competency 1389", Audience: "employees", Duration: 17}

// L1390 catalog extension
var lesson_1390 = LessonIndex{Code: "L1390", Title: "advanced module 1390", Objective: "Apply competency 1390", Audience: "employees", Duration: 18}

// L1391 catalog extension
var lesson_1391 = LessonIndex{Code: "L1391", Title: "advanced module 1391", Objective: "Apply competency 1391", Audience: "employees", Duration: 19}

// L1392 catalog extension
var lesson_1392 = LessonIndex{Code: "L1392", Title: "advanced module 1392", Objective: "Apply competency 1392", Audience: "employees", Duration: 20}

// L1393 catalog extension
var lesson_1393 = LessonIndex{Code: "L1393", Title: "advanced module 1393", Objective: "Apply competency 1393", Audience: "employees", Duration: 21}

// L1394 catalog extension
var lesson_1394 = LessonIndex{Code: "L1394", Title: "advanced module 1394", Objective: "Apply competency 1394", Audience: "employees", Duration: 22}

// L1395 catalog extension
var lesson_1395 = LessonIndex{Code: "L1395", Title: "advanced module 1395", Objective: "Apply competency 1395", Audience: "employees", Duration: 23}

// L1396 catalog extension
var lesson_1396 = LessonIndex{Code: "L1396", Title: "advanced module 1396", Objective: "Apply competency 1396", Audience: "employees", Duration: 24}

// L1397 catalog extension
var lesson_1397 = LessonIndex{Code: "L1397", Title: "advanced module 1397", Objective: "Apply competency 1397", Audience: "employees", Duration: 25}

// L1398 catalog extension
var lesson_1398 = LessonIndex{Code: "L1398", Title: "advanced module 1398", Objective: "Apply competency 1398", Audience: "employees", Duration: 26}

// L1399 catalog extension
var lesson_1399 = LessonIndex{Code: "L1399", Title: "advanced module 1399", Objective: "Apply competency 1399", Audience: "employees", Duration: 27}

// L1400 catalog extension
var lesson_1400 = LessonIndex{Code: "L1400", Title: "advanced module 1400", Objective: "Apply competency 1400", Audience: "employees", Duration: 8}

// L1401 catalog extension
var lesson_1401 = LessonIndex{Code: "L1401", Title: "advanced module 1401", Objective: "Apply competency 1401", Audience: "employees", Duration: 9}

// L1402 catalog extension
var lesson_1402 = LessonIndex{Code: "L1402", Title: "advanced module 1402", Objective: "Apply competency 1402", Audience: "employees", Duration: 10}

// L1403 catalog extension
var lesson_1403 = LessonIndex{Code: "L1403", Title: "advanced module 1403", Objective: "Apply competency 1403", Audience: "employees", Duration: 11}

// L1404 catalog extension
var lesson_1404 = LessonIndex{Code: "L1404", Title: "advanced module 1404", Objective: "Apply competency 1404", Audience: "employees", Duration: 12}

// L1405 catalog extension
var lesson_1405 = LessonIndex{Code: "L1405", Title: "advanced module 1405", Objective: "Apply competency 1405", Audience: "employees", Duration: 13}

// L1406 catalog extension
var lesson_1406 = LessonIndex{Code: "L1406", Title: "advanced module 1406", Objective: "Apply competency 1406", Audience: "employees", Duration: 14}

// L1407 catalog extension
var lesson_1407 = LessonIndex{Code: "L1407", Title: "advanced module 1407", Objective: "Apply competency 1407", Audience: "employees", Duration: 15}

// L1408 catalog extension
var lesson_1408 = LessonIndex{Code: "L1408", Title: "advanced module 1408", Objective: "Apply competency 1408", Audience: "employees", Duration: 16}

// L1409 catalog extension
var lesson_1409 = LessonIndex{Code: "L1409", Title: "advanced module 1409", Objective: "Apply competency 1409", Audience: "employees", Duration: 17}

// L1410 catalog extension
var lesson_1410 = LessonIndex{Code: "L1410", Title: "advanced module 1410", Objective: "Apply competency 1410", Audience: "employees", Duration: 18}

// L1411 catalog extension
var lesson_1411 = LessonIndex{Code: "L1411", Title: "advanced module 1411", Objective: "Apply competency 1411", Audience: "employees", Duration: 19}

// L1412 catalog extension
var lesson_1412 = LessonIndex{Code: "L1412", Title: "advanced module 1412", Objective: "Apply competency 1412", Audience: "employees", Duration: 20}

// L1413 catalog extension
var lesson_1413 = LessonIndex{Code: "L1413", Title: "advanced module 1413", Objective: "Apply competency 1413", Audience: "employees", Duration: 21}

// L1414 catalog extension
var lesson_1414 = LessonIndex{Code: "L1414", Title: "advanced module 1414", Objective: "Apply competency 1414", Audience: "employees", Duration: 22}

// L1415 catalog extension
var lesson_1415 = LessonIndex{Code: "L1415", Title: "advanced module 1415", Objective: "Apply competency 1415", Audience: "employees", Duration: 23}

// L1416 catalog extension
var lesson_1416 = LessonIndex{Code: "L1416", Title: "advanced module 1416", Objective: "Apply competency 1416", Audience: "employees", Duration: 24}

// L1417 catalog extension
var lesson_1417 = LessonIndex{Code: "L1417", Title: "advanced module 1417", Objective: "Apply competency 1417", Audience: "employees", Duration: 25}

// L1418 catalog extension
var lesson_1418 = LessonIndex{Code: "L1418", Title: "advanced module 1418", Objective: "Apply competency 1418", Audience: "employees", Duration: 26}

// L1419 catalog extension
var lesson_1419 = LessonIndex{Code: "L1419", Title: "advanced module 1419", Objective: "Apply competency 1419", Audience: "employees", Duration: 27}

// L1420 catalog extension
var lesson_1420 = LessonIndex{Code: "L1420", Title: "advanced module 1420", Objective: "Apply competency 1420", Audience: "employees", Duration: 8}

// L1421 catalog extension
var lesson_1421 = LessonIndex{Code: "L1421", Title: "advanced module 1421", Objective: "Apply competency 1421", Audience: "employees", Duration: 9}

// L1422 catalog extension
var lesson_1422 = LessonIndex{Code: "L1422", Title: "advanced module 1422", Objective: "Apply competency 1422", Audience: "employees", Duration: 10}

// L1423 catalog extension
var lesson_1423 = LessonIndex{Code: "L1423", Title: "advanced module 1423", Objective: "Apply competency 1423", Audience: "employees", Duration: 11}

// L1424 catalog extension
var lesson_1424 = LessonIndex{Code: "L1424", Title: "advanced module 1424", Objective: "Apply competency 1424", Audience: "employees", Duration: 12}

// L1425 catalog extension
var lesson_1425 = LessonIndex{Code: "L1425", Title: "advanced module 1425", Objective: "Apply competency 1425", Audience: "employees", Duration: 13}

// L1426 catalog extension
var lesson_1426 = LessonIndex{Code: "L1426", Title: "advanced module 1426", Objective: "Apply competency 1426", Audience: "employees", Duration: 14}

// L1427 catalog extension
var lesson_1427 = LessonIndex{Code: "L1427", Title: "advanced module 1427", Objective: "Apply competency 1427", Audience: "employees", Duration: 15}

// L1428 catalog extension
var lesson_1428 = LessonIndex{Code: "L1428", Title: "advanced module 1428", Objective: "Apply competency 1428", Audience: "employees", Duration: 16}

// L1429 catalog extension
var lesson_1429 = LessonIndex{Code: "L1429", Title: "advanced module 1429", Objective: "Apply competency 1429", Audience: "employees", Duration: 17}

// L1430 catalog extension
var lesson_1430 = LessonIndex{Code: "L1430", Title: "advanced module 1430", Objective: "Apply competency 1430", Audience: "employees", Duration: 18}

// L1431 catalog extension
var lesson_1431 = LessonIndex{Code: "L1431", Title: "advanced module 1431", Objective: "Apply competency 1431", Audience: "employees", Duration: 19}

// L1432 catalog extension
var lesson_1432 = LessonIndex{Code: "L1432", Title: "advanced module 1432", Objective: "Apply competency 1432", Audience: "employees", Duration: 20}

// L1433 catalog extension
var lesson_1433 = LessonIndex{Code: "L1433", Title: "advanced module 1433", Objective: "Apply competency 1433", Audience: "employees", Duration: 21}

// L1434 catalog extension
var lesson_1434 = LessonIndex{Code: "L1434", Title: "advanced module 1434", Objective: "Apply competency 1434", Audience: "employees", Duration: 22}

// L1435 catalog extension
var lesson_1435 = LessonIndex{Code: "L1435", Title: "advanced module 1435", Objective: "Apply competency 1435", Audience: "employees", Duration: 23}

// L1436 catalog extension
var lesson_1436 = LessonIndex{Code: "L1436", Title: "advanced module 1436", Objective: "Apply competency 1436", Audience: "employees", Duration: 24}

// L1437 catalog extension
var lesson_1437 = LessonIndex{Code: "L1437", Title: "advanced module 1437", Objective: "Apply competency 1437", Audience: "employees", Duration: 25}

// L1438 catalog extension
var lesson_1438 = LessonIndex{Code: "L1438", Title: "advanced module 1438", Objective: "Apply competency 1438", Audience: "employees", Duration: 26}

// L1439 catalog extension
var lesson_1439 = LessonIndex{Code: "L1439", Title: "advanced module 1439", Objective: "Apply competency 1439", Audience: "employees", Duration: 27}

// L1440 catalog extension
var lesson_1440 = LessonIndex{Code: "L1440", Title: "advanced module 1440", Objective: "Apply competency 1440", Audience: "employees", Duration: 8}

// L1441 catalog extension
var lesson_1441 = LessonIndex{Code: "L1441", Title: "advanced module 1441", Objective: "Apply competency 1441", Audience: "employees", Duration: 9}

// L1442 catalog extension
var lesson_1442 = LessonIndex{Code: "L1442", Title: "advanced module 1442", Objective: "Apply competency 1442", Audience: "employees", Duration: 10}

// L1443 catalog extension
var lesson_1443 = LessonIndex{Code: "L1443", Title: "advanced module 1443", Objective: "Apply competency 1443", Audience: "employees", Duration: 11}

// L1444 catalog extension
var lesson_1444 = LessonIndex{Code: "L1444", Title: "advanced module 1444", Objective: "Apply competency 1444", Audience: "employees", Duration: 12}

// L1445 catalog extension
var lesson_1445 = LessonIndex{Code: "L1445", Title: "advanced module 1445", Objective: "Apply competency 1445", Audience: "employees", Duration: 13}

// L1446 catalog extension
var lesson_1446 = LessonIndex{Code: "L1446", Title: "advanced module 1446", Objective: "Apply competency 1446", Audience: "employees", Duration: 14}

// L1447 catalog extension
var lesson_1447 = LessonIndex{Code: "L1447", Title: "advanced module 1447", Objective: "Apply competency 1447", Audience: "employees", Duration: 15}

// L1448 catalog extension
var lesson_1448 = LessonIndex{Code: "L1448", Title: "advanced module 1448", Objective: "Apply competency 1448", Audience: "employees", Duration: 16}

// L1449 catalog extension
var lesson_1449 = LessonIndex{Code: "L1449", Title: "advanced module 1449", Objective: "Apply competency 1449", Audience: "employees", Duration: 17}

// L1450 catalog extension
var lesson_1450 = LessonIndex{Code: "L1450", Title: "advanced module 1450", Objective: "Apply competency 1450", Audience: "employees", Duration: 18}
