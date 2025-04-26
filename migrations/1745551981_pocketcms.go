package migrations

import (
	"fmt"

	"github.com/2g8/pocketcms/core"
	"github.com/2g8/pocketcms/tools/types"
)


func init() {
	core.SystemMigrations.Register(func(txApp core.App) error {
		if err := createMembersCollection(txApp); err != nil {
			return fmt.Errorf("members error: %w", err)
		}
		if err := createPostsCollection(txApp); err != nil {
			return fmt.Errorf("posts error: %w", err)
		}
		if err := createPostsAuthorsCollection(txApp); err != nil {
			return fmt.Errorf("posts_authors error: %w", err)
		}
		if err := createTagsCollection(txApp); err != nil {
			return fmt.Errorf("tags error: %w", err)
		}
		if err := createPostsMetaCollection(txApp); err != nil {
			return fmt.Errorf("posts_meta error: %w", err)
		}
		if err := createPostsTagsCollection(txApp); err != nil {
			return fmt.Errorf("posts_tags error: %w", err)
		}
		if err := createPostsRevisionsCollection(txApp); err != nil {
			return fmt.Errorf("post_revisions error: %w", err)
		}
		if err := createNewslettersCollection(txApp); err != nil {
			return fmt.Errorf("newsletters error: %w", err)
		}
		if err := createSubscriptionsCollection(txApp); err != nil {
			return fmt.Errorf("subscriptions error: %w", err)
		}

		return nil
	}, func(txApp core.App) error {
		tables := []string{
			"members",
			"posts",
			"posts_authors",
			"tags",
			"posts_meta",
			"posts_tags",
			"post_revisions",
			"newsletters",
			"subscriptions",
		}

		for _, name := range tables {
			if _, err := txApp.DB().DropTable(name).Execute(); err != nil {
				return err
			}
		}

		return nil
	})
}


func createMembersCollection(txApp core.App) error {
	members := core.NewAuthCollection("members", "_pc_members_auth_")

	ownerRule := "id = @request.auth.id"
	members.ListRule = types.Pointer(ownerRule)
	members.ViewRule = types.Pointer(ownerRule)
	members.CreateRule = types.Pointer("")
	members.UpdateRule = types.Pointer(ownerRule)
	members.DeleteRule = types.Pointer(ownerRule)

	members.Fields.Add(&core.TextField{
		Name: "transient_id",
		Max:  254,
		Required: true,
	})
	members.Fields.Add(&core.EmailField{
		Name: "email",
		Required: true,
	})
	members.Fields.Add(&core.TextField{
		Name: "username",
		Max:  254,
	})
	members.Fields.Add(&core.PasswordField{
		Name:     "password",
		System:   true,
		Hidden:   true,
		Required: true,
		Cost:     8,
	})
	members.Fields.Add(&core.TextField{
		Name: "status",
		Max:  50,
		Required: true,
		AutogeneratePattern: "free",
	})
	members.Fields.Add(&core.TextField{
		Name: "name",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "slug",
		Max:  254,
	})
	members.Fields.Add(&core.FileField{
		Name:      "avatar",
		MaxSelect: 1,
		MimeTypes: []string{"image/jpeg", "image/png", "image/svg+xml", "image/gif", "image/webp"},
	})
	members.Fields.Add(&core.TextField{
		Name: "cover_image",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "bio",
	})
	members.Fields.Add(&core.TextField{
		Name: "website",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "location",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "facebook",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "twitter",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "linkedin",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "accessibility",
	})
	members.Fields.Add(&core.TextField{
		Name: "locale",
		Max:  6,
	})
	members.Fields.Add(&core.TextField{
		Name: "visibility",
		Max:  50,
		AutogeneratePattern: "public",
	})
	members.Fields.Add(&core.TextField{
		Name: "meta_title",
		Max:  500,
	})
	members.Fields.Add(&core.TextField{
		Name: "meta_description",
		Max:  500,
	})
	members.Fields.Add(&core.TextField{
		Name: "tour",
	})
	members.Fields.Add(&core.DateField{
		Name: "last_seen",
	})
	members.Fields.Add(&core.TextField{
		Name: "expertise",
		Max:  254,
	})
	members.Fields.Add(&core.TextField{
		Name: "note",
		Max:  2000,
	})
	members.Fields.Add(&core.TextField{
		Name: "geolocation",
		Max:  254,
	})
	members.Fields.Add(&core.BoolField{
		Name: "email_disabled",
		Required: true,
	})
	members.Fields.Add(&core.AutodateField{
		Name:     "created_at",
		OnCreate: true,
	})
	members.Fields.Add(&core.TextField{
		Name: "created_by",
		Required: true,
		Max: 50,
	})
	members.Fields.Add(&core.AutodateField{
		Name:     "updated_at",
		OnCreate: true,
		OnUpdate: true,
	})
	members.Fields.Add(&core.TextField{
		Name: "updated_by",
		Max: 50,
	})

	members.AddIndex("members_transient_id_unique", true, "transient_id", "")
	members.AddIndex("members_slug_unique", true, "slug", "")
	members.AddIndex("members_username_unique", true, "username", "")

	members.OAuth2.MappedFields.Username = "username"
	members.OAuth2.MappedFields.Name = "name"
	members.OAuth2.MappedFields.AvatarURL = "avatar"

	return txApp.Save(members)
}

func createPostsCollection(txApp core.App) error {
	posts := core.NewBaseCollection("posts")

	posts.Fields.Add(&core.TextField{
		Name: "member_id",
		Max:  36,
		Required: true,
	})
	posts.Fields.Add(&core.TextField{
		Name: "title",
		Max:  500,
		Required: true,
	})
	posts.Fields.Add(&core.TextField{
		Name: "slug",
		Max:  254,
		Required: true,
	})
	posts.Fields.Add(&core.TextField{
		Name: "mobiledoc",
	})
	posts.Fields.Add(&core.TextField{
		Name: "lexical",
	})
	posts.Fields.Add(&core.TextField{
		Name: "html",
	})
	posts.Fields.Add(&core.TextField{
		Name: "comment_id",
		Max:  36,
	})
	posts.Fields.Add(&core.TextField{
		Name: "plaintext",
	})
	posts.Fields.Add(&core.TextField{
		Name: "feature_image",
		Max:  254,
	})
	posts.Fields.Add(&core.BoolField{
		Name: "featured",
		Required: true,
	})
	posts.Fields.Add(&core.TextField{
		Name: "type",
		Max:  50,
		Required: true,
		AutogeneratePattern: "post",
	})
	posts.Fields.Add(&core.TextField{
		Name: "status",
		Max:  50,
		Required: true,
		AutogeneratePattern: "draft",
	})
	posts.Fields.Add(&core.TextField{
		Name: "locale",
		Max:  6,
	})
	posts.Fields.Add(&core.TextField{
		Name: "visibility",
		Max:  50,
		Required: true,
		AutogeneratePattern: "public",
	})
	posts.Fields.Add(&core.TextField{
		Name: "email_recipient_filter",
		Required: true,
	})
	posts.Fields.Add(&core.AutodateField{
		Name:     "created_at",
		OnCreate: true,
	})
	posts.Fields.Add(&core.TextField{
		Name: "created_by",
		Max:  36,
		Required: true,
	})
	posts.Fields.Add(&core.AutodateField{
		Name:     "updated_at",
		OnCreate: true,
		OnUpdate: true,
	})
	posts.Fields.Add(&core.TextField{
		Name: "updated_by",
		Max:  36,
	})
	posts.Fields.Add(&core.DateField{
		Name: "published_at",
	})
	posts.Fields.Add(&core.TextField{
		Name: "published_by",
		Max:  36,
	})
	posts.Fields.Add(&core.TextField{
		Name: "custom_excerpt",
		Max:  2000,
	})
	posts.Fields.Add(&core.TextField{
		Name: "codeinjection_head",
	})
	posts.Fields.Add(&core.TextField{
		Name: "codeinjection_foot",
	})
	posts.Fields.Add(&core.TextField{
		Name: "custom_template",
		Max:  100,
	})
	posts.Fields.Add(&core.TextField{
		Name: "canonical_url",
	})
	posts.Fields.Add(&core.TextField{
		Name: "newsletter_id",
		Max:  36,
	})
	posts.Fields.Add(&core.BoolField{
		Name: "show_title_and_feature_image",
		Required: true,
	})

	posts.AddIndex("posts_slug_type_unique", true, "slug,type", "")
	posts.AddIndex("posts_updated_at_index", false, "updated_at", "")
	posts.AddIndex("posts_published_at_index", false, "published_at", "")
	posts.AddIndex("posts_newsletter_id_foreign", false, "newsletter_id", "")
	posts.AddIndex("posts_type_status_updated_at_index", false, "type,status,updated_at", "")

	return txApp.Save(posts)
}

func createPostsAuthorsCollection(txApp core.App) error {
	postsAuthors := core.NewBaseCollection("posts_authors")

	postsAuthors.Fields.Add(&core.TextField{
		Name: "post_id",
		Max:  36,
	})
	postsAuthors.Fields.Add(&core.TextField{
		Name: "author_id",
		Max:  36,
	})
	postsAuthors.Fields.Add(&core.NumberField{
		Name: "sort_order",
	})

	// 添加索引
	postsAuthors.AddIndex("posts_authors_post_id_foreign", false, "post_id", "")
	postsAuthors.AddIndex("posts_authors_author_id_foreign", false, "author_id", "")

	
	return txApp.Save(postsAuthors)
}

func createTagsCollection(txApp core.App) error {
	tags := core.NewBaseCollection("tags")

	tags.Fields.Add(&core.TextField{
		Name: "name",
		Max:  254,
		Required: true,
	})
	tags.Fields.Add(&core.TextField{
		Name: "slug",
		Max:  254,
		Required: true,
	})
	tags.Fields.Add(&core.TextField{
		Name: "description",
	})
	tags.Fields.Add(&core.TextField{
		Name: "feature_image",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "parent_id",
		Max:  36,
	})
	tags.Fields.Add(&core.TextField{
		Name: "visibility",
		Max:  50,
		AutogeneratePattern: "public",
	})
	tags.Fields.Add(&core.TextField{
		Name: "og_image",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "og_title",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "og_description",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "twitter_image",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "twitter_title",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "twitter_description",
		Max:  500,
	})
	tags.Fields.Add(&core.TextField{
		Name: "meta_title",
		Max:  500,
	})
	tags.Fields.Add(&core.TextField{
		Name: "meta_description",
		Max:  500,
	})
	tags.Fields.Add(&core.TextField{
		Name: "codeinjection_head",
	})
	tags.Fields.Add(&core.TextField{
		Name: "codeinjection_foot",
	})
	tags.Fields.Add(&core.TextField{
		Name: "canonical_url",
		Max:  254,
	})
	tags.Fields.Add(&core.TextField{
		Name: "accent_color",
		Max:  50,
	})
	tags.Fields.Add(&core.AutodateField{
		Name:     "created_at",
		OnCreate: true,
	})
	tags.Fields.Add(&core.TextField{
		Name: "created_by",
		Max: 36,
	})
	tags.Fields.Add(&core.AutodateField{
		Name:     "updated_at",
		OnCreate: true,
		OnUpdate: true,
	})
	tags.Fields.Add(&core.TextField{
		Name: "updated_by",
		Max: 36,
	})

	// 添加唯一索引
	tags.AddIndex("tags_slug_unique", true, "slug", "")

	return txApp.Save(tags)
}

func createPostsMetaCollection(txApp core.App) error {
	postsMeta := core.NewBaseCollection("posts_meta")

	postsMeta.Fields.Add(&core.TextField{
		Name: "post_id",
		Max:  36,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "og_image",
		Max:  254,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "og_title",
		Max:  300,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "og_description",
		Max:  500,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "twitter_image",
		Max:  254,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "twitter_title",
		Max:  300,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "twitter_description",
		Max:  500,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "meta_title",
		Max:  500,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "meta_description",
		Max:  500,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "email_subject",
		Max:  300,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "frontmatter",
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "feature_image_alt",
		Max:  254,
	})
	postsMeta.Fields.Add(&core.TextField{
		Name: "feature_image_caption",
	})
	postsMeta.Fields.Add(&core.BoolField{
		Name: "email_only",
		Required: true,
	})

	// 添加唯一索引
	postsMeta.AddIndex("posts_meta_post_id_unique", true, "post_id", "")

	return txApp.Save(postsMeta)
}

func createPostsTagsCollection(txApp core.App) error {
	postsTags := core.NewBaseCollection("posts_tags")

	postsTags.Fields.Add(&core.TextField{
		Name: "post_id",
		Max:  36,
	})
	postsTags.Fields.Add(&core.TextField{
		Name: "tag_id",
		Max:  36,
	})
	postsTags.Fields.Add(&core.NumberField{
		Name: "sort_order",
	})

	// 添加索引
	postsTags.AddIndex("posts_tags_tag_id_foreign", false, "tag_id", "")
	postsTags.AddIndex("posts_tags_post_id_tag_id_index", false, "post_id,tag_id", "")

	return txApp.Save(postsTags)
}

func createNewslettersCollection(txApp core.App) error {
	newsletters := core.NewBaseCollection("newsletters")

	newsletters.Fields.Add(&core.TextField{
		Name: "member_id",
		Max:  36,
		Required: true,
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "name",
		Max:  254,
		Required: true,
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "description",
		Max:  2000,
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "feedback_enabled",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "slug",
		Max:  254,
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "sender_name",
		Max:  254,
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "sender_email",
		Max:  254,
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "sender_reply_to",
		Max:  254,
		AutogeneratePattern: "newsletter",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "status",
		Max:  50,
		AutogeneratePattern: "active",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "visibility",
		Max:  50,
		AutogeneratePattern: "members",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "subscribe_on_signup",
	})
	newsletters.Fields.Add(&core.NumberField{
		Name: "sort_order",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "header_image",
		Max:  254,
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_header_icon",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_header_title",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_excerpt",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "title_font_category",
		Max:  254,
		AutogeneratePattern: "sans_serif",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "title_alignment",
		Max:  254,
		AutogeneratePattern: "center",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_feature_image",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "body_font_category",
		Max:  254,
		AutogeneratePattern: "sans_serif",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "footer_content",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_badge",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_header_name",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_post_title_section",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_comment_cta",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_subscription_details",
	})
	newsletters.Fields.Add(&core.BoolField{
		Name: "show_latest_posts",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "background_color",
		Max:  50,
		AutogeneratePattern: "light",
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "border_color",
		Max:  50,
	})
	newsletters.Fields.Add(&core.TextField{
		Name: "title_color",
		Max:  50,
	})
	newsletters.Fields.Add(&core.AutodateField{
		Name:     "created_at",
		OnCreate: true,
	})
	newsletters.Fields.Add(&core.AutodateField{
		Name:     "updated_at",
		OnCreate: true,
		OnUpdate: true,
	})

	// 添加唯一索引
	newsletters.AddIndex("newsletters_member_id_unique", true, "member_id", "")
	newsletters.AddIndex("newsletters_name_unique", true, "name", "")
	newsletters.AddIndex("newsletters_slug_unique", true, "slug", "")

	return txApp.Save(newsletters)
}

func createSubscriptionsCollection(txApp core.App) error {
	subscriptions := core.NewBaseCollection("subscriptions")

	subscriptions.Fields.Add(&core.TextField{
		Name: "member_id",
		Max:  36,
		Required: true,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "type",
		Max:  50,
		Required: true,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "status",
		Max:  50,
		Required: true,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "tier_id",
		Max:  36,
		Required: true,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "cadence",
		Max:  50,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "currency",
		Max:  50,
	})
	subscriptions.Fields.Add(&core.NumberField{
		Name: "amount",
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "payment_provider",
		Max:  50,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "payment_subscription_url",
		Max:  2000,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "payment_user_url",
		Max:  2000,
	})
	subscriptions.Fields.Add(&core.TextField{
		Name: "offer_id",
		Max:  36,
	})
	subscriptions.Fields.Add(&core.DateField{
		Name: "expires_at",
	})
	subscriptions.Fields.Add(&core.AutodateField{
		Name:     "created_at",
		OnCreate: true,
	})
	subscriptions.Fields.Add(&core.AutodateField{
		Name:     "updated_at",
		OnCreate: true,
		OnUpdate: true,
	})

	// 添加索引
	subscriptions.AddIndex("subscriptions_member_id_foreign", false, "member_id", "")
	subscriptions.AddIndex("subscriptions_tier_id_foreign", false, "tier_id", "")
	subscriptions.AddIndex("subscriptions_offer_id_foreign", false, "offer_id", "")

	return txApp.Save(subscriptions)
}

func createPostsRevisionsCollection(txApp core.App) error {
	postsRevisions := core.NewBaseCollection("post_revisions")

	postsRevisions.Fields.Add(&core.TextField{
		Name: "post_id",
		Max:  36,
		Required: true,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "lexical",
	})
	postsRevisions.Fields.Add(&core.DateField{
		Name: "created_at",
		Required: true,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "author_id",
		Max:  36,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "title",
		Max:  500,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "post_status",
		Max:  50,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "reason",
		Max:  50,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "feature_image",
		Max:  254,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "feature_image_alt",
		Max:  254,
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "feature_image_caption",
	})
	postsRevisions.Fields.Add(&core.TextField{
		Name: "custom_excerpt",
		Max:  2000,
	})

	// 添加索引
	postsRevisions.AddIndex("post_revisions_post_id_index", false, "post_id", "")
	postsRevisions.AddIndex("post_revs_author_id_foreign", false, "author_id", "")

	return txApp.Save(postsRevisions)
}

