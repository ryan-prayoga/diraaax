# diraaax Database Schema v2

This document defines the database schema.

Database: PostgreSQL

---

# users

id
display_name
nickname
slug
created_at

---

# timeline_events

id
title
description
event_date
image_url
created_by
created_at

---

# memories

id
title
description
image_url
memory_date
created_by
created_at

---

# bucket_list

id
title
description
is_done
target_date
created_by
created_at

---

# love_capsules

id
title
message
open_date
created_by
created_at

---

# voice_notes

id
audio_url
created_by
created_at

---

# daily_moods

id
user_id
mood
date

---

# memory_locations

id
title
description
lat
lng
image_url
created_at

---

# love_reasons

id
message
created_by
created_at
