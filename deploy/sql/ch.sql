CREATE TABLE default.app_static_analysis_detail
(
    `app_id`                String,
    `sdk_names`             Array(String),
    `version_code`          String,
    `version_name`          String,
    `developer`             String,
    `is_google_listing`     Bool,
    `file_path`             String,
    `created_at`            DateTime64(3)
) ENGINE = MergeTree
ORDER BY tuple()
SETTINGS index_granularity = 8192;

-- `default`.android_app_detail definition
CREATE TABLE default.android_app_detail
(
    `title`                 String,
    `label`                 String,
    `icon_url`              String,
    `package_name`          String,
    `version_code`          String,
    `version_name`          String,
    `sign`                  Array(String),
    `review_stars`          Float64,
    `description`           String,
    `description_short`     String,
    `whatsnew`              String,
    `asset_usability`       String,
    `developer`             String,
    `is_show_comment_score` Bool,
    `comment_score1`        String,
    `comment_score2`        String,
    `comment_score3`        String,
    `comment_score4`        String,
    `comment_score5`        String,
    `comment_total`         String,
    `comment_score_total`   String,
    `comment_score_stars`   Float64,
    `price`                 String,
    `in_app_products`       String,
    `introduction`          String,
    `category_name`         String,
    `update_date`           DateTime,
    `create_date`           DateTime,
    `is_free`               Bool,
    `tags`                  Array(String),
    `sha1`                  String,
    `size`                  String,
    `is_google_listing`     Bool,
    `download_count`        String,
    `version_id`            String,
    `app_id`                String,
    `native_code`           Array(String),
    `apk_type`              UInt8,
    `real_package_name`     String,
    `sdk_version`           String,
    `target_sdk_version`    String,
    `created_at`            DateTime
) ENGINE = ReplacingMergeTree
PRIMARY KEY real_package_name
ORDER BY real_package_name
SETTINGS index_granularity = 8192;